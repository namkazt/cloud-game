From golang:1.12

RUN apt-get update

RUN apt-get install pkg-config libvpx-dev libopus-dev libopusfile-dev -y

RUN mkdir -p /cloud-game
VOLUME /home/thanh/Desktop/Code/cloud-game/games:/cloud-game/games
COPY . /cloud-game/
WORKDIR /cloud-game

# Install server dependencies
RUN go install ./cmd

EXPOSE 8000
