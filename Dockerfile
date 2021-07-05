FROM alpine
# DB를 쓰는 경우 Timezone 데이터가 필요한데 alpine에는 기본적으로 존재하지 않음
RUN apk add tzdata
WORKDIR /khumu
# root directory가 아니라 build한 output임.
# github action으로 빌드함.
COPY khumu-club /khumu/khumu-club
ENV KHUMU_HOME /khumu
ENV KHUMU_ENVIRONMENT DEV
CMD ["./khumu-club"]