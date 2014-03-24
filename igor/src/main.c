#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ncurses.h>

static int WIN_H = 0;
static int WIN_W = 0;

char * drawh( int n, const char * s );
char * drawh( int n, const char * s ) {
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
  //ncurses startup
  initscr();

  //ncurses defaults
  start_color();
  curs_set(0);
  getmaxyx(stdscr,WIN_H,WIN_W);
  raw();
  keypad(stdscr, TRUE);
  noecho();

  //errors
  if(has_colors() == FALSE) { endwin();
    printf("Your terminal does not support color\n");
    exit(1);
  }

  if(WIN_H < 25 || WIN_W < 50) {
    printf("Your terminal is WAY too small\n"); 
    exit(1);
  }
  int i;

  init_pair(1, COLOR_BLACK, COLOR_WHITE);
  init_pair(2, COLOR_BLUE, COLOR_WHITE);
  init_pair(3, COLOR_RED, COLOR_WHITE);
  init_pair(99, COLOR_BLUE, COLOR_WHITE);

  attron(COLOR_PAIR(1));
  char * background = drawh(WIN_W, " ");
  for(i=0;i<WIN_H;++i) {
    move(i,0);
    printw(background);
  } 
  free(background);
  attroff(COLOR_PAIR(1));

  //display header & footer
  char * header = drawh(WIN_W, ".");
  attron(COLOR_PAIR(1));
  move(0,0);
  printw(header);
  move(WIN_H-1,0);
  printw(header);
  attroff(COLOR_PAIR(1));

  attron(COLOR_PAIR(2));
  move(1,0);
  printw(header);
  move(WIN_H-2,0);
  printw(header);
  attroff(COLOR_PAIR(2));

  attron(COLOR_PAIR(3));
  move(2,0);
  printw(header);
  move(WIN_H-3,0);
  printw(header);
  attroff(COLOR_PAIR(3));

  attron(COLOR_PAIR(99));
  move(4,4);
  printw("@");
  attroff(COLOR_PAIR(99));


  free(header);

  move((WIN_H/2)-1,0);
  printw("Press any key to begin");
  //user input
  int ch = getch();
  if(ch == KEY_F(1)) {
    printw("F1 Key pressed");
  } else {
    move(WIN_H/2,0);
    printw("The pressed key is ");
    attron(A_BOLD);
    printw("%c", ch);
    attroff(A_BOLD);
  }

  move(WIN_H-2,0);
  printw("** press any key to exit **");
  getch();
  //refresh();
  endwin();
  return 0;
}
