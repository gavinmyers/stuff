all: main.c 
	gcc main.c -v -lncurses -L -Wall -Wunused-function -Wunused-value -Wnonnull -Wvolatile-register-var -Wunused-variable -Werror -Wextra -Wformat=2 -Wno-format-nonliteral -Wshadow -Wpointer-arith -Wcast-qual -Wmissing-prototypes -Wno-missing-braces -std=gnu89 -D_GNU_SOURCE -O2 -o igor
test:
	clang main.c -cc1 -analyze -analyzer-checker=core -analyzer-output=text -v -lncurses -L -Wall -Wextra -Wformat=2 -Wno-format-nonliteral -Wshadow -Wpointer-arith -Wcast-qual -Wmissing-prototypes -Wno-missing-braces -std=gnu89 -D_GNU_SOURCE -O2 -o igor.html
