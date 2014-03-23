#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ncurses.h>

static int WIN_H = 0;
static int WIN_W = 0;

char * drawline( int n, const char * s );
char * drawline( int n, const char * s ) {
  size_t slen = strlen(s);
  char * dest = malloc(n*slen+1);
 
  int i; char * p;
  for ( i=0, p = dest; i < n; ++i, p += slen ) {
    memcpy(p, s, slen);
  }
  *p = '\0';
  return dest;
}

int main() { 
  int ch;
  initscr();
  if(has_colors() == FALSE)
  { endwin();
    printf("Your terminal does not support color\n");
    exit(1);
  }
  start_color();
  curs_set(0);
  getmaxyx(stdscr,WIN_H,WIN_W);
  raw();
  keypad(stdscr, TRUE);
  noecho();

  attron(COLOR_PAIR(1));
  char * header = drawline(WIN_W, "#");
  move(0,0);
  init_pair(1, COLOR_RED, COLOR_BLACK);
  printw(header);
  move(WIN_H-1,0);
  printw(header);
  free(header);
  attroff(COLOR_PAIR(1));

  ch = getch();
  if(ch == KEY_F(1)) {
    printw("F1 Key pressed");
  } else {
    move(WIN_H/2,0);
    printw("The pressed key is ");
    attron(A_BOLD);
    printw("%c", ch);
    attroff(A_BOLD);
  }
  refresh();
  getch();
  endwin();
  return 0;
}
