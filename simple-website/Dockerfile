FROM debian:buster-slim as builder

RUN apt-get update

RUN apt-get install -y \
	git \
	cmake \
	make \
	zlib1g-dev \
	build-essential \
	libssl-dev

RUN git clone --recurse-submodules --depth=1 https://github.com/h2o/h2o.git
RUN cd h2o && mkdir -p build && cd build && cmake .. && make && make install


FROM debian:buster-slim

COPY --from=builder /usr/local/bin/h2o /usr/local/bin
COPY --from=builder /usr/local/share/h2o /usr/local/share/h2o

RUN apt-get update

RUN apt-get install -y libssl-dev

COPY h2o.conf /etc/h2o/h2o.conf
COPY ./www /var/www/localhost

EXPOSE 80

RUN h2o -v && h2o --conf /etc/h2o/h2o.conf --test

CMD h2o --conf /etc/h2o/h2o.conf

