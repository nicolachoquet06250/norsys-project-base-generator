package configFiles

import (
	. "npbg/helpers"
	technosEnum "npbg/technos/enum"
)

type treeElement = map[string]string
type tree = map[string]treeElement
type treeGeneration struct{}

func (g treeGeneration) javaScript(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.JavaScript] = treeElement{
		"site" + Slash() + "index.html": `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta name="viewport"
				  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Project Base Generator</title>
		<script type="module">
			import myFunction from "./script.js"
			myFunction()
		</script>
	</head>
	<body>
		<h1>Welcome to <b>Project Base Generator</b></h1>
		<p></p>
	</body>
</html>`,

		"site" + Slash() + "script.js": `
export default function () {
	document.querySelector('p').innerHTML = 'ça fontionne !!'
}`,

		"Dockerfile": `
FROM httpd:2.4
WORKDIR /usr/local/apache2/htdocs/
COPY ./site/ /usr/local/apache2/htdocs/`,

		"build.sh": "sudo docker build −t static−image .",
		"run.sh":   "sudo docker run −p 80:80 −−name static−image−1 static−image",
	}

	return &g
}

func (g treeGeneration) react15(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.React15] = treeElement{
		"Dockerfile": `
FROM node:latest
WORKDIR /app
COPY package.json /app
COPY yarn.lock /app
RUN yarn
COPY . /app
EXPOSE 3000
CMD ["yarn","start"]!`,

		"build.sh": "docker image build . -t mon-app:dev",
		"run.sh":   "docker container run -d -p 3000:3000 -v /home/arthur/Documents/GWT/mon-app/src:/app/src mon-app:dev",
	}

	return &g
}

func (g treeGeneration) react16(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.React16] = treeElement{
		"Dockerfile": `
FROM node:latest
WORKDIR /app
COPY package.json /app
COPY yarn.lock /app
RUN yarn
COPY . /app
EXPOSE 3000
CMD ["yarn","start"]!`,

		"build.sh": "docker image build . -t mon-app:dev",
		"run.sh":   "docker container run -d -p 3000:3000 -v /home/arthur/Documents/GWT/mon-app/src:/app/src mon-app:dev",
	}

	return &g
}

func (g treeGeneration) react17(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.React17] = treeElement{
		"Dockerfile": `
FROM node:latest
WORKDIR /app
COPY package.json /app
COPY yarn.lock /app
RUN yarn
COPY . /app
EXPOSE 3000
CMD ["yarn","start"]!`,

		"build.sh": "docker image build . -t mon-app:dev",
		"run.sh":   "docker container run -d -p 3000:3000 -v /home/arthur/Documents/GWT/mon-app/src:/app/src mon-app:dev",
	}

	return &g
}

func (g treeGeneration) react18(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.React18] = treeElement{
		"Dockerfile": `
FROM node:latest
WORKDIR /app
COPY package.json /app
COPY yarn.lock /app
RUN yarn
COPY . /app
EXPOSE 3000
CMD ["yarn","start"]!`,

		"build.sh": "docker image build . -t mon-app:dev",
		"run.sh":   "docker container run -d -p 3000:3000 -v /home/arthur/Documents/GWT/mon-app/src:/app/src mon-app:dev",
	}

	return &g
}

func (g treeGeneration) vue2(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.Vue2] = treeElement{
		"Dockerfile": `
FROM node:lts-alpine
# installe un simple serveur http pour servir un contenu statique
RUN npm install -g http-server
# définit le dossier 'app' comme dossier de travail
WORKDIR /app
# copie 'package.json' et 'package-lock.json' (si disponible)
COPY package*.json ./
# installe les dépendances du projet
RUN npm install
# copie les fichiers et dossiers du projet dans le dossier de travail (par exemple : le dossier 'app')
COPY . .
# construit l'app pour la production en la minifiant
RUN npm run build
EXPOSE 8080
CMD [ "http-server", "dist" ]
`,

		"build.sh": "docker build -t my-vuejs-app .",
		"run.sh":   "docker run -it -p 8080:8080 --rm --name my-vuejs-app-1 my-vuejs-app",
	}

	return &g
}

func (g treeGeneration) vue3(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.Vue3] = treeElement{
		"Dockerfile": `
FROM node:lts-alpine
# installe un simple serveur http pour servir un contenu statique
RUN npm install -g http-server
# définit le dossier 'app' comme dossier de travail
WORKDIR /app
# copie 'package.json' et 'package-lock.json' (si disponible)
COPY package*.json ./
# installe les dépendances du projet
RUN npm install
# copie les fichiers et dossiers du projet dans le dossier de travail (par exemple : le dossier 'app')
COPY . .
# construit l'app pour la production en la minifiant
RUN npm run build
EXPOSE 8080
CMD [ "http-server", "dist" ]
`,

		"build.sh": "docker build -t my-vuejs-app .",
		"run.sh":   "docker run -it -p 8080:8080 --rm --name my-vuejs-app-1 my-vuejs-app",
	}

	return &g
}

func (g treeGeneration) angular(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.Angular] = treeElement{
		"Dockerfile": `
FROM node:11.15.0-stretch
RUN npm install -g @angular/cli && ng config -g cli.packageManager yarn
WORKDIR /app
version: '3'
services:
  node:
    build: .
    ports:
      - "4200:4200"
    volumes:
      - ".:/app"
    tty: true
	command: [sh, -c, cd my-project && ng serve --host=0.0.0.0 --poll=2000]
`,

		"build.sh": "docker-compose up -d",
	}

	return &g
}

func (g treeGeneration) php(configFiles *tree) *treeGeneration {
	(*configFiles)[technosEnum.PHP] = treeElement{
		"docker-compose.yml": `
version: '3.3'
services:
  php_apache:
    build:
      context: ./php
      dockerfile: Dockerfile
    container_name: php8_0_2_apache
    depends_on:
      - mysql
    volumes:
      - ./{{.ProjectName}}:/var/www/html/
    ports:
      - 8000:80
  mysql:
    container_name: mysql_
    command: "--innodb_use_native_aio=0"
    volumes:
      - "./db_data:/var/lib/mysql"
      - "./mysql_conf:/etc/mysql/conf.d"
    environment:
      MYSQL_ROOT_PASSWORD: 12345
      MYSQL_DATABASE: docker_{{.ProjectName}}
    ports:
      - 3306:3306
`,
		"php" + Slash() + "Dockerfile": `
FROM php:8.0.2-apache
RUN apt-get update && apt-get upgrade -y
EXPOSE 80
`,
		"run.sh":                               "docker-compose up --build -d",
		"{{.ProjectName}}" + Slash() + ".void": "*",
	}

	return &g
}
