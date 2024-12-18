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

Create a file as shown below. Replace the values With your own Jenkins server configuration. 
```
{ 
  "host" : "http://localhost:8080",
  "user": "my-user", 
  "token": "my-token"
}
```
And then run the command below. If you receive a success message with the Jenkins server info it means that you are good to go!
```
jctl info
```

## Examples

### List Builds
To list all the build from a especified job run the command:
```
jctl list builds --jobId "my-job"
```
Where jobId is the full project name of the job, so in case the job is created inside a folder, add the whole path:
```
jctl list builds --jobId "my-root-folder/my-subfolder/my-job"
```
Additionally you can add a the field `jobId` in the `config.json` file to take the value from there by default, althought it is not mandatory.

### Create a new Build

To create and run a new build run the following:
```
jctl create build --jobId "my-job" --params params.json
```
Where the `params.json` contains a json file with a list of all the parametres used to run the job. You can ignore this flag if the pipeline does not have any parameters.
The file must have a key/value structure, with only string values are supported, as Jenkins only support this type of parametres. An example of a `params.json` file is shown below:
```
{
  "STRING_PARAM": "mystring",
  "BOOLEAN_PARAM": "false",
  "NUMBER_PARAM": "20"
}

```





