
variable "bucket_name" {
  description = "The bucket name of the Huaweicloud OBS"
  default     = "tf_bucket_demo"
}

variable "obs_address" {
  description = "The address (/endpoint) of the Huaweicloud OBS"
  default     = "obs.cn-north-4.myhuaweicloud.com"
}

variable "object_path" {
  description = "The objct path where the file is located in the Huaweicloud OBS"
  default     = "for_download/"
}

variable "object_name" {
  description = "The objct name of the Huaweicloud OBS"
  default     = "README.md"
}

# If the file does not exist in the current folder,
# you need to fill in the absolute path of the folder where the user file is located.
variable "local_file_path" {
  description = "The local path where the file is located"
  default     = "README.md"
}

variable "function_name" {
  description = "The function name of the Huaweicloud FunctionGraph"
  default     = "tf_function_demo"
}

variable "trigger_name" {
  description = "The trigger name of the Huaweicloud FunctionGraph"
  default     = "tf_trigger_demo"
}
