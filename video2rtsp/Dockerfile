FROM jrottenberg/ffmpeg:3.4-alpine

COPY ffserver.conf /etc/ffserver.conf
COPY run.sh /run.sh

ENTRYPOINT ["sh"]
CMD ["/run.sh"]