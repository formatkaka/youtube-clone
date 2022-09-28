# Building Youtube clone

Instructions to run


    git clone https://github.com/formatkaka/youtube-clone/

## Run with docker

    // Build
    docker build -t yt-db-1 -f Dockerfile.db . 
    docker build -t yt-sv-1 -f Dockerfile.server .
    docker build -t yt-cl-1 -f Dockerfile.client .  

    // Run
    docker run -d -p 5431:5432 yt-db-1
    docker run -d -p 8080:8080 yt-sv-1
    docker run -d -p 5173:5173 yt-cl-2

-----
## Run on system directly
### Set up DB

    Install psql locally
    Create db
        name - youtube
        user - postgres
        password - postgres
    Run create.sql -> add-data.sql 

### Server Setup

    Install Go 
    cd server/
    go mod download
    go run main.go

## Client Setup

    Install node v16 (use nvm)
    cd client/
    npm i
    npm run dev -- --open


## Hosting

[Hosting Diagram](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&layers=1&nav=1&title=Youtube-clone.drawio#RxZbbjtsgEIafxlJ7sZJt1m5ymeO20rZJ60rbWxKIjRZDSkji9Ok7xNjxKa1X2s1eGf6BYfhmwDhokmYPCm%2BTr5JQ7vguyRw0dXzf84cD%2BBjllCsD5OdCrBixgy5CxP5QK7pW3TNCd7WBWkqu2bYurqUQdK1rGlZKHuvDNpLXV93imLaEaI15W31iRCd2F4F70T9TFifFyp5rLSkuBlthl2AijxUJzRw0UVLqvJVmE8oNvIJLPm9%2BxVoGpqjQfSaQ0a9vIg7C0P3ymwfZYjiKnu6slwPme7thG6w%2BFQSU3AtCjRPXQeNjwjSNtnhtrEfIOWiJTjn0PGhad1Rpml2N0yt3D2VDZUq1OsEQO6EEZitmaLvHC34%2FtFpSRT%2BwIrYpj0vXFyrQsGBeAMnvgBRyWHZM2AGasWlGB8o1fWa6sMFSFTN8cWpgidXOfD7MZwYzVcDqY4s5BAUFTm%2FCGwV13p7fBu55HcA%2FvRVv1OL9IDkWMWjRGdi78vJRg5fbwcu9Ja%2F7Fq9l9P3xfSGFDUgdp%2Fi2RRVcPcSsckbPwqoQpguYMf2xWD7OflaO9ao5AbSKkwZ2gKjrfHdayWc6kVwqUIQUJicbxnlDwpzFArprgA5Fj8YmJQz%2BTSNrSBkh%2FFpC6%2Ff2a1wUXj2nYc%2B6R2%2BV0rDXvbwcTWbR1Uu5ObysAdWR5H9PnUPuIBp4r0izW0FM%2FKl5X5h6kOrc6uu9UUT%2FObF4t80fPxuWmZS%2FRrpDv%2FlfCPod4VJ8QcKhe3kJnW2V9ySa%2FQU%3D)


# Project Development

1. [Home Page](https://github.com/formatkaka/youtube-clone/milestone/1)
2. [Video Watch Page](https://github.com/formatkaka/youtube-clone/milestone/3)
3. [Studio Page to upload videos](https://github.com/formatkaka/youtube-clone/milestone/5)
4. [Search Page](https://github.com/formatkaka/youtube-clone/milestone/2)
5. [Channel Page](https://github.com/formatkaka/youtube-clone/milestone/4)
