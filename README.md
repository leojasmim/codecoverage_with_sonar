# codecoverage_with_sonar
Este repositório tem como objetivo exemplificar como utilizar a ferramenta de CodeCoverage no Sonar com aplicações em GO

O SonarQube é uma plataforma de código aberto desenvolvida pela SonarSource para inspeção contínua da qualidade do código, para executar revisões automáticas com análise estática do código para detectar bugs, code smells e vulnerabilidades de segurança e está disponível em mais de 20 linguagens de programação. Sua integração é muito utilizada no processo de CI. Para maiores informações acesse o site do [SonarQube](https://www.sonarqube.org/).

Este repositório tem por objetivo mostrar como integrar a análise da cobertura de código (_code coverage_) no Sonar para aplicações GO. 

## A Aplicação

A aplicação é uma função bem simples que calcula se um número inteiro é par ou impar. Criamos o arquivo _main_test.go_ no qual temos o teste unitário dos 2 fluxos existentes na função. Vamos integrar lá no Sonar utilizando a análise de _code coverage_.

## Pré-Requisitos para esse tutorial

1. Ter uma aplicação em Go com testes
2. Docker
3. JRE (_Java Runtime Environment_)

## Como fazer

### Gerando o arquivo de CodeCoverage

Antes de utilizar o Sonar precisamos rodar os nossos testes e gerar o arquivo de code coverage. Neste exemplo salvaremos esse arquivo na raiz do repositório com o nome codecoverage.file. Para gerar o arquivo executaremos o seguinte comando:

`go test ./... -short -coverprofile=coverage.file`

Para englobar todos os testes do projeto utilizamos o `./...`

![codecoverage_with_sonar](https://raw.githubusercontent.com/leojasmim/codecoverage_with_sonar/main/docs/create_coverage_file.png)

### Adicionando as configurações sonar no Projeto

Agora iremos adicionar a raiz do projeto o arquivo de configuração para o SonarQube chamado **sonar-project.properties**. Este arquivo é responsável por informar ao SonarQube quais os códigos deverão ser analisados, nome e versão do projeto e onde localizar o arquivo de _code coverage_. Temos um exemplo deste arquivo de configuração:

``` YAML
sonar.projectKey=ljasmim_coverage_with_sonar
sonar.projectName=covarage_with_sonar

sonar.sourceEncoding=UTF-8

sonar.sources=.
sonar.exclusions=**/*_test.go,**/vendor/**

sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.test.exclusions=**/vendor/**

sonar.go.coverage.reportPaths=**/coverage.file

sonar.scm.disabled=true
```

Sobre alguns desses parâmetros podemos destacar:

* `sonar.source` : Indica a localização do código a ser analisado. Neste exemplo temos a raiz do projeto
* `sonar.exclusions` : Arquivos que deverão ser ignorados na análise. No exemplo não será analisado qualquer arquivo de teste da aplicação, nem arquivos do go vendor.
* `sonar.test.inclusions` : Arquivos de teste a serem executados.
* `sonar.test.exclusions` : Arquivos de teste a serem ignorados.
* `sonar.go.coverage.reportPaths` : Indica a localização do arquivo de _code coverage_
* `sonar.scm.disabled` : Desabilita a função _Source Control Manager_. Dentre as diversas funções o SCM faz **a detecção automática de novos códigos**, logo o checagem poderá ser ignorada no caso de repositórios já existentes no Sonar. Maiores informações sobre _SCM Integration_ na [documentação](https://docs.sonarqube.org/latest/analysis/scm-integration/).

### Executando e configurando o SonarQube

Para analisar o código com o SonarQube precisamos rodar o serviço do Sonar. Isso poderá ser feito com o docker executando o comando

`docker run --name sonarqube -p 9000:9000 sonarqube`

Além disso, para utilizar a ferramenta _sonar-scanner_ provida pelo SonarQube precisaremos do JRE na máquina. Você poderá fazer o download da ferramenta neste [link](https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/)

Uma vez que o docker estiver rodando devemos acessar o serviço (Administration/Security/Users) e criar um token de acesso:

![token_generation](https://raw.githubusercontent.com/leojasmim/codecoverage_with_sonar/main/docs/create_token.png)

Com esse token podemos utilizar o _sonar-scanner_ passando a _project-key_ e o _token_ de acesso:

`sonar-scanner.bat -D"sonar.projectKey=ljasmim_coverage_with_sonar" -D"sonar.sources=." -D"sonar.host.url=http://localhost:9000" -D"sonar.login={MEU_TOKEN_AQUI}"`

Abaixo temos o log da execução do _sonar-scanner_ :

```
C:\Users\ljasmin\go\src\codecoverage_with_sonar>sonar-scanner.bat -D"sonar.projectKey=ljasmim_coverage_with_sonar" -D"sonar.sources=." -D"sonar.host.url=http://localhost:9000" -D"sonar.login=fe84d629f23d10c9dafff06a56796289ccfd9967"
INFO: Scanner configuration file: C:\Users\ljasmin\workspaces\sonar-scanner-4.6.0.2311-windows\bin\..\conf\sonar-scanner.properties
INFO: Project root configuration file: C:\Users\ljasmin\go\src\codecoverage_with_sonar\sonar-project.properties
INFO: SonarScanner 4.6.0.2311
INFO: Java 11.0.3 AdoptOpenJDK (64-bit)
INFO: Windows 10 10.0 amd64
INFO: User cache: C:\Users\ljasmin\.sonar\cache
INFO: Scanner configuration file: C:\Users\ljasmin\workspaces\sonar-scanner-4.6.0.2311-windows\bin\..\conf\sonar-scanner.properties
INFO: Project root configuration file: C:\Users\ljasmin\go\src\codecoverage_with_sonar\sonar-project.properties
INFO: Analyzing on SonarQube server 8.7.0
INFO: Default locale: "pt_BR", source code encoding: "UTF-8"
INFO: Load global settings
INFO: Load global settings (done) | time=114ms
INFO: Server id: BF41A1F2-AXgSM68yuLS1tDA0ZehF
INFO: User cache: C:\Users\ljasmin\.sonar\cache
INFO: Load/download plugins
INFO: Load plugins index
INFO: Load plugins index (done) | time=85ms
INFO: Load/download plugins (done) | time=194ms
INFO: Process project properties
INFO: Process project properties (done) | time=10ms
INFO: Execute project builders
INFO: Execute project builders (done) | time=3ms
INFO: Project key: ljasmim_coverage_with_sonar
INFO: Base dir: C:\Users\ljasmin\go\src\codecoverage_with_sonar
INFO: Working dir: C:\Users\ljasmin\go\src\codecoverage_with_sonar\.scannerwork
INFO: Load project settings for component key: 'ljasmim_coverage_with_sonar'
INFO: Load project settings for component key: 'ljasmim_coverage_with_sonar' (done) | time=52ms
INFO: Load quality profiles
INFO: Load quality profiles (done) | time=71ms
INFO: Load active rules
INFO: Load active rules (done) | time=2357ms
INFO: Indexing files...
INFO: Project configuration:
INFO:   Excluded sources: **/*_test.go, **/*_test.go
INFO:   Included tests: **/*_test.go
INFO:   Excluded tests: **/vendor/**
INFO: 7 files indexed
INFO: 7 files ignored because of inclusion/exclusion patterns
INFO: Quality profile for go: Sonar way
INFO: ------------- Run sensors on module covarage_with_sonar
INFO: Load metrics repository
INFO: Load metrics repository (done) | time=51ms
INFO: Sensor CSS Rules [cssfamily]
INFO: No CSS, PHP, HTML or VueJS files are found in the project. CSS analysis is skipped.
INFO: Sensor CSS Rules [cssfamily] (done) | time=4ms
INFO: Sensor Code Quality and Security for Go [go]
INFO: 1 source files to be analyzed
INFO: Load project repositories
INFO: Load project repositories (done) | time=43ms
INFO: Sensor Code Quality and Security for Go [go] (done) | time=605ms
INFO: 1/1 source files have been analyzed
INFO: Sensor Go Cover sensor for Go coverage [go]
INFO: Load coverage report from 'C:\Users\ljasmin\go\src\codecoverage_with_sonar\coverage.file'
INFO: Sensor Go Cover sensor for Go coverage [go] (done) | time=59ms
INFO: Sensor JaCoCo XML Report Importer [jacoco]
INFO: 'sonar.coverage.jacoco.xmlReportPaths' is not defined. Using default locations: target/site/jacoco/jacoco.xml,target/site/jacoco-it/jacoco.xml,build/reports/jacoco/test/jacocoTestReport.xml
INFO: No report imported, no coverage information will be imported by JaCoCo XML Report Importer
INFO: Sensor JaCoCo XML Report Importer [jacoco] (done) | time=12ms
INFO: Sensor C# Properties [csharp]
INFO: Sensor C# Properties [csharp] (done) | time=11ms
INFO: Sensor JavaXmlSensor [java]
INFO: Sensor JavaXmlSensor [java] (done) | time=6ms
INFO: Sensor HTML [web]
INFO: Sensor HTML [web] (done) | time=6ms
INFO: Sensor VB.NET Properties [vbnet]
INFO: Sensor VB.NET Properties [vbnet] (done) | time=3ms
INFO: ------------- Run sensors on project
INFO: Sensor Zero Coverage Sensor
INFO: Sensor Zero Coverage Sensor (done) | time=2ms
INFO: SCM Publisher is disabled
INFO: CPD Executor Calculating CPD for 1 file
INFO: CPD Executor CPD calculation finished (done) | time=15ms
INFO: Analysis report generated in 122ms, dir size=92 KB
INFO: Analysis report compressed in 60ms, zip size=12 KB
INFO: Analysis report uploaded in 32ms
INFO: ANALYSIS SUCCESSFUL, you can browse http://localhost:9000/dashboard?id=ljasmim_coverage_with_sonar
INFO: Note that you will be able to access the updated dashboard once the server has processed the submitted analysis report
INFO: More about the report processing at http://localhost:9000/api/ce/task?id=AXgYeYqG9b475blKcHG9
INFO: Analysis total time: 5.232 s
INFO: ------------------------------------------------------------------------
INFO: EXECUTION SUCCESS
INFO: ------------------------------------------------------------------------
INFO: Total time: 7.135s
INFO: Final Memory: 11M/47M
INFO: ------------------------------------------------------------------------
```

Acessando o navegador conforme indicado no log temos o nosso projeto no SonarQube com o _code coverage_:

![dash01](https://raw.githubusercontent.com/leojasmim/codecoverage_with_sonar/main/docs/project_dash_01.png)

![dash02](https://raw.githubusercontent.com/leojasmim/codecoverage_with_sonar/main/docs/project_dash_02.png)

O _sonar_scanner_ irá incluir uma pasta `.scannerwork` na estrutura do projeto. É recomendado adicioná-la ao gitignore do projeto para evitar o versionamento.

## Referência

STAMER, Jan. Go for SonarQube. Disponível em https://medium.com/red6-es/go-for-sonarqube-ffff5b74f33a. Acesso em 08/03/2021.
