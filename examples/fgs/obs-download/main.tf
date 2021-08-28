
resource "huaweicloud_obs_bucket" "test" {
  bucket = var.bucket_name
  acl    = "private"
}

resource "huaweicloud_obs_bucket_object" "test" {
  bucket = huaweicloud_obs_bucket.test.bucket
  key    = format("%s%s", var.object_path, var.object_name)
  source = var.local_file_path
}

resource "huaweicloud_identity_agency" "test" {
  name                   = "obs_agency"
  delegated_service_name = "op_svc_cff"

  project_role {
    project = "MOS"
    roles = [
      "OBS OperateAccess",
    ]
  }
  domain_roles = [
    "OBS OperateAccess",
    "IAM ReadOnlyAccess",
  ]
}

resource "huaweicloud_fgs_function" "test" {
  name        = var.function_name
  app         = "default"
  agency      = huaweicloud_identity_agency.test.name
  description = "Download file from OBS bucket"
  handler     = "index.handler"
  depend_list = [
    "a12cea7e-8a52-4cf5-8ab5-5956d1cf7325",
  ]
  memory_size = 256
  timeout     = 15
  runtime     = "Python2.7"
  code_type   = "inline"
  user_data   = jsonencode({"srcBucket"=var.bucket_name,"srcObjPath"=var.object_path,"srcObjName"=var.object_name,"obsAddress"=var.obs_address})
  func_code   = <<EOF
# -*- coding: utf-8 -*-
# When running this sample code to access OBS, you must specify an agency with global service access permissions (or at least with OBS access permissions).
from obs import ObsClient #Require public dependency:esdk_obs_python-3.x
import sys
import os

current_file_path = os.path.dirname(os.path.realpath(__file__))
# Adds the current path to search paths to import third-party libraries.
sys.path.append(current_file_path)

TEMP_ROOT_PATH = "/tmp/"       # Downloads a file from OBS to this directory.

# Handler of the function
def handler (event, context):
    logger = context.getLogger()                    # Obtains a log instance.

    srcBucket = context.getUserData('srcBucket')    # Enter the name of the bucket where the actual file to be downloaded is stored.
    srcObjName = context.getUserData('srcObjName')  # Enter the name of the actual file to be downloaded, for example, file.txt.
    srcObjPath = context.getUserData('srcObjPath')  # Enter the directoty of the actual file to be downloaded, for example, for_download/.

    if srcBucket is None or srcObjName is None:
        logger.error("Please set environment variables srcBucket and srcObjName.")
        return ("Please set environment variables srcBucket and srcObjName.")

    # You are advised to use the log instance provided by FunctionGraph to debug or print messages and not to use the native print function.
    logger.info( "*** srcBucketName: " + srcBucket)
    logger.info("*** srcObjName:" + srcObjName)

    # Obtains a temporary AK and SK to access OBS. An agency is required to access IAM.
    ak = context.getAccessKey()
    sk = context.getSecretKey()
    if ak == "" or sk == "":
        logger.error("Failed to access OBS because no temporary AK, SK, or token has been obtained. Please set an agency.")
        return ("Failed to access OBS because no temporary AK, SK, or token has been obtained. Please set an agency.")

    obs_address = context.getUserData('obsAddress')      # Domain name of the OBS service. Use the default value.
    if obs_address is None:
        obs_address = 'obs.cn-north-1.myhuaweicloud.com'
    # Build a client.
    client = GetObsClient(obs_address, ak, sk)
    # Downloads an uploaded file from OBS.
    status = GetObject(client, srcBucket, srcObjPath, srcObjName)
    if (status == 200 or status == 201):
        image_size = cal_image_size(srcObjName)
        output = "the object you download from OBS is " + str(image_size) + " KB"
        logger.info(output)
        return output
    else:
        logger.error("download file from OBS failed")
        return ("download file from OBS failed")


# Build a obs client to get the obs object.
def GetObsClient(obsAddr, ak, sk):
    client = ObsClient(access_key_id=ak, secret_access_key=sk,server=obsAddr)
    return client


# Downloads a file from OBS to a local directory.
def GetObject(client, bucketName, objPath, objName):
    resp = client.getObject(bucketName, objPath+objName, downloadPath=TEMP_ROOT_PATH+objName)
    print('*** GetObject resp: ', resp)
    return (int(resp.status))


def cal_image_size(fileName):
    fileNamePath = TEMP_ROOT_PATH + fileName
    # Calculates the size of a file in KB.
    size = os.path.getsize(fileNamePath) / 1024
    return size

EOF
}

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"
  status       = "ACTIVE"

  timer {
    name          = var.trigger_name
    schedule_type = "Rate"
    schedule      = "3d"
  }
}
