# jenkins-cli

Simple jenkins go CLI which uses the [gojenkins](https://www.github.com/bndr/gojenkins) API library under the hood.

Current supported features:
- List jobs, builds and artifacts.
- Download artifacts.
- Get build info.
- Show build logs.

## Install jctl

To install the tool golang and make should be previously installed in your system. Download the repository and run the command:
```
make
```
The command line `jctl` tool will be installed in your PATH. This tool only runs in linux and MacOS.

## Run jctl

The tool will look for credentials in the `~/.jctl/config.json` by default.

Create a file as shown below. Replace the values with your own Jenkins server configuration. The `host`,`user`i and `token` values are mandatory.
```
{ 
  "host" : "http://localhost:8080",
  "user": "my-user", 
  "token": "my-token"
  "job": "my-root-folder/my-subfolder/my-job"
}
```
And `job`, which is an optional field, corresponds to the full project name of the job. In case the job is created inside a folder, add the whole path. Finally run the command below. 
```
jctl info
```
If you receive a success message with the Jenkins server info it means that you are good to go!

## Examples

### List Builds
To list all the build from the specified run the command:
```
jctl list builds
```
Additionally you can specify the job from the command line.
```
jctl list builds --jobId "my-root-folder/my-subfolder/my-job"
```

### Create a new Build

To create and run a new build run the following:
```
jctl create build --params params.env
```
Where the `params.env` consists of a `.env` file with all the parameters as variables. You can ignore this flag if the pipeline does not need any parameters. Example of a `.env` file is shown below:
```
"STRING_PARAM"="mystring"
"BOOLEAN_PARAM"=false
"NUMBER_PARAM"=20
```

### View the logs

In order to view the logs of a specific build, run the following command:
```
jctl logs --build <build-number>
```

Additionally, you can also check the logs of the last build:
```
jctl logs --latest --follow
```

To `follow` flag, in this case, will follow the logs of a running build.





