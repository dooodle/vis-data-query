FROM scratch
ADD server /
CMD ["/server","-serve"]