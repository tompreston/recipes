vpath %.yaml recipes 

POSTS_DIR := content/posts
POSTS := \
	${POSTS_DIR}/spaghetti-carbonara.md \
	${POSTS_DIR}/pork-ramen.md \
	${POSTS_DIR}/martini.md \
	${POSTS_DIR}/rice.md

.PHONY: all clean
all: ${POSTS}

${POSTS_DIR}/%.md: %.yaml
	nom $^ > $@

clean:
	rm ${POSTS}
