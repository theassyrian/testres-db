# testres-db

is a tool and library to import test results to a single database. It allows to
analyze your efficiency of testing, improve test execution and get better
vizualization of testing in your project. Queries can be executed with a
standard `sqlite` client or using [Jupyter Notebook](https://jupyter.org/).
Below there are some examples of SQL queries:

```sql
~$ sqlite testres-db.sqlite
SQLite version 2.8.17
Enter ".help" for instructions
sqlite> SELECT * FROM testrun TR JOIN test T ON TR.test_id = T.test_id JOIN status ST ON TR.status_id = ST.status_id WHERE ST.name = 'PASSED';
sqlite> SELECT * FROM testrun TR JOIN test T ON TR.test_id = T.test_id JOIN suite S ON T.suite_id = S.suite_id JOIN report R ON TR.report_id = R.report_id JOIN status ST ON TR.status_id = ST.status_id WHERE ST.name = 'PASSED';
```

## Building

```
$ go get ./...
$ go test -v ./...
$ cd cmd/testres-db && go build .
```

## Documentation

There are several ways to obtain test results from external services:
1. First one is getting them via API, but only few services supports such methods in API
1. Second one is getting file with test report in standard format, like JUnit,
[TestAnythingProtocol](https://testanything.org/) or SubUnit. Moreover
`testres-db` supports [Go test JSON format](https://golang.org/cmd/test2json/)
and [GTester output format](https://wiki.gnome.org/Projects/GLib/GTester).

First of all you should make sure test reports is available after every build
on continuous integration.

Secondly you should describe configuration file in YAML format which contains
information about your project.

```yaml
projects:
- name: criu
  backends:
   - type: jenkins
     branch: master
     pipeline: ""
     base: https://ci.openvz.org/
     username: anonymous
     secret: anonymous
```

By default `testres-db` uses `testres-db.yaml` as configuration file and
`testres.sqlite` as an SQLite database file. You can change default values
using command-line options.

### Azure DevOps

Generate access token in
[profile](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops&tabs=preview-page).

### Circle CI

Circle CI
[supports](https://circleci.com/docs/2.0/configuration-reference/#store_test_results)
storing test results in two formats: JUnit XML and Cucumber JSON. Access token
is not required to query info for public projects, to access to a private
projects follow steps in
[documentation](https://circleci.com/docs/api/#add-an-api-token) to obtain API
token.

### Cirrus CI

Cirrus CI allows to [store test report
files](https://cirrus-ci.org/guide/writing-tasks/#artifacts-instruction) in
JUnit formats. Access to Cirrus CI API may require access token. You can
generate it in [profile](https://cirrus-ci.com/settings/profile/).
Note: by default [Cirrus CI persists caches and logs for 90 days](https://cirrus-ci.org/faq/).

### GitLab

Gitlab supports publishing JUnit reports. Follow [instructions in
documentation](https://docs.gitlab.com/ee/ci/junit_test_reports.html) to setup
it up. In general case to enable the JUnit reports in merge requests, you need
to add `artifacts:reports:junit` in `.gitlab-ci.yml`, and specify the path(s)
of the generated test reports. Access to Gitlab API may require access token
with `api` capability. To generate access token please refer to
[documentation](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html).
`expire_in`
[allows](https://docs.gitlab.com/ce/ci/yaml/README.html#artifactsexpire_in) you
to specify how long artifacts should live before they expire and therefore
deleted, counting from the time they are uploaded and stored on GitLab. If the
expiry time is not defined, it defaults to the instance wide setting (30 days
by default, forever on GitLab.com).

### Jenkins

Jenkins supports publishing test results in JUnit format, but not all Jenkins
instances supports it. To check this you may point your browser to a link
`<Jenkins URL>/job/<Job Name>/lastCompletedBuild/testReport/api/xml`. Browser
will show XML document in case of success. Access to Jenkins API may require
access token, please refer to
[documentation](https://jenkins.io/doc/book/using/using-credentials/) to obtain
it.

### TeamCity CI

Access to TeamCity API requires access token, it can be generate in
[Settings](https://travis-ci.org/account/preferences).
Currently, supported API version is TeamCity 2018.1.3 or up.

### Travis CI

Unfortunately Travis CI doesn't support publishing of test results.  Access to
Travis API requires access token, it can be generate in
[Settings](https://travis-ci.org/account/preferences). Username value is not
required in config.

Sample of config:

```yaml
- type: travisci
  branch: criu-dev
  pipeline: checkpoint-restore/criu
  base: https://travis-ci.org/
  secret: Imyr-a2jvoNc1qy5nnoOvA
```

## Contribution

There are numerous awesome [Continuous Integration
services](https://github.com/ligurio/awesome-ci) which are not integrated with
`testres-db`. Feel free to make a patches and bring support for them.

## Authors

Developed with passion by [Sergey Bronnikov](https://bronevichok.ru/) and great
open source [contributors](https://github.com/ligurio/testres-db/contributors).
