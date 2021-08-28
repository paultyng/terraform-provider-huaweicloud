## Example: Basic NAT gateway and dnat rule

This example provisions:

- a ECS instance:
  > It can mount one or more data disks.
  > Using CentOS 7.3 64bit as the operating system image.
- a VPC and subnet instance:
  > VPC provides an isolated cloud service environment.
- a basic NAT gateway.
- a EIP.
- a snat rule:
  > By binding external subnet IP, SNAT function can realize multiple virtual machines across availability zones to share external subnet IP and access external data center or other VPCs.
- a dnat rule:
  > DNAT rule specifies port 8080 as the port for Tomcat to provide external services, which can be used to build nginx services.
  > DNAT function is bound with EIP, and EIP is shared across VPC by binding IP mapping, which provides services for the Internet.
  > After building completed, the nginx service can be accessed through the open port 8080 of DNAT.
- a remote-exec provisioner:
  > Automatic sequential execution of multiple terminal commands.
