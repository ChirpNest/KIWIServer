# KIWIServer

### Install Go on WSL
When using Windows, activate the Windows subsystem for Linux (WSL) (link) and install the Linux distribution “Ubuntu 20.04 LTS” (link) and set it up (don’t forget to run sudo apt update && sudo apt upgrade).

Download the file “go1.14.4.linux-amd64.tar.gz” from https://golang.org/dl/.
Extract the file to “/usr/local” with the following command:
```sh
$	sudo tar -C /usr/local -xzf /mnt/c/Users/my_user/path_to_file/go1.14.4.linux-amd64.tar.gz
```
Add the following rows to the file “~/.profile” (e.g. use the command sudo nano ~/.profile):
```sh
# add GOPATH and GOBIN and add GO to the PATH environment variable
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export PATH=$PATH:/usr/local/go/bin
```
Leave the command line with the command exit and open it again (so a logon is performed).
Just as a note: The official install guide for Go can be found at https://golang.org/doc/install.

### Clone the Repository and get started
Clone the repository “kiwi-server” to the local file system (in WSL accessible at /mnt/c/…) using your favourite git tool (note that the git tool should not replace line feeds “LF” with carriage return and line feeds “CRLF”) or the command line as follows:
```sh
$	git clone URL_TO_GIT_REPOSITORY
```
Now try to build kiwi-server with the following command (use sudo apt install make if make is not yet installed):
```sh
$	make build
```
Now install the development requirements with the following command:
```sh
$	make dev-requirements
```
Now create a snapshot with the following command:
```sh
$	make snapshot
```
The snapshot which runs on the LORIX One (Linux, ARM 7) can be transmitted to the LORIX One using the following command (replace IP_ADDRESS with the IP address of the LORIX One):
```sh
$	make copy-to-lorix host=IP_ADDRESS
```
After transmitting it to the LORIX One, it can be started on the LORIX One when connected via SSH (see “Configure ChirpNest via SSH” on page 111) with the following command:
 the LORIX One (Linux, ARM 7) can be transmitted to the LORIX One using the following command (replace IP_ADDRESS with the IP address of the LORIX One):
```sh
$	~/kiwi-server
```

### Edit the Source Code
To edit the source code we recommend installing Visual Studio Code (https://code.visualstudio.com/). 
Under Windows: use the extension “Remote - WSL” (https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) and work remote on the WSL.
Also install the “Go” extension (on WSL when using Windows) (https://marketplace.visualstudio.com/items?itemName=golang.Go).


### Generate API Code from .proto Files
When the definition of the API is changed in “measurement.proto” and/or “device.proto”, the server files (“.pb.go” and “.pb.gw.go”) can be generated automatically by performing the following steps.
Install the following packages:

```sh
$	sudo apt install protobuf-compiler libprotobuf-dev
```
Change the directory as follows:
```sh
$	cd apidefinition/go/external
```
Run the following commands:
```sh
$	make requirements
$	make api-external
```
Now all “.pb.go” and “.pb.gw.go” were generated.

### Unit Tests

To run the unit tests, a correctly configured 
```sh
$	make test
```
If the error message exec: "gcc": executable file not found in $PATH appears, run the following command:
```sh
$	sudo apt install build-essential
```
make test should now run the tests, but they fail because a correctly configured PostgreSQL database is required.
To install PostgreSQL locally (recommended), run the following command:
```sh
$	sudo apt install postgresql postgresql-contrib
```
Then, start the PostgreSQL instance with the following command (might be necessary again after rebooting the system):
```sh
$	sudo service postgresql start
```
Follow the instruction in the code when searching for “unit test howto: create required database”. When this is done, the unit tests should pass when triggered with make test.

### Add new Version to ChirpNest

How to create a ChirpNest Image is explained in “VII.III Manual: Create ChirpNest Yocto Image for LORIX One”. This section explains how the source code of the “chirpstack-gateway-os” repository can be updated so a new version of KIWI Server is installed.

First, create a snapshot using make snapshot (as explained above). Then, the file “kiwi_server_v0.0.0-next_Linux_armv7.tar.gz” from the “dist” folder needs to be uploaded somewhere where it can be directly accessed through an URL.

Now edit the following file in the “chirpstack-gateway-os”:
```sh
layers/chirpstack/meta-chirpstack/recipes-chirpstack/kiwi-server/kiwi-server_BETA.bb
```
Replace the URL so it points to the newly uploaded file (search for “.tar.gz”). Also replace the SHA-256 checksum (search for “sha256sum”). The SHA-256 checksum can be determined by the following command (replace PATH_TO_FILE with the path to the file ending with “.tar.gz”):

```sh
$	sha256sum PATH_TO_FILE
```
That’s it, now the image can be created as explained in “VII.III Manual: Create ChirpNest Yocto Image for LORIX One”.











