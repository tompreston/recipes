vpath %.yaml recipes 

POSTS_DIR := content/posts
POSTS := \
	${POSTS_DIR}/spaghetti-carbonara.md \
	${POSTS_DIR}/pork-ramen.md \
	${POSTS_DIR}/martini.md \
	${POSTS_DIR}/rice.md

.PHONY: all test install server clean
all: ${POSTS}

${POSTS_DIR}/%.md: %.yaml
	nom $^ > $@

test:
	make -C tools/nom test

install:
	make -C tools/nom install

server:
	hugo server -D

clean:
	rm ${POSTS}
