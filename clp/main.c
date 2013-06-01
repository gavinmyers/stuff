#include <stdio.h>
#include <stdlib.h>
#include <GLUT/glut.h> 
//#include <GL/glut.h>//Drawing funciton

#define KEY_ESC 27
#define WINDOW_POS_X 100
#define WINDOW_POS_Y 100
#define DEF_SCREEN_WIDTH 1224 
#define DEF_SCREEN_HEIGHT 700 

static int width = 0;
static int height = 0;

static GLubyte *buffer = NULL;
static int fullscreen = 0;

static void
keyboard(unsigned char c, int x, int y)
{
  switch (c) {
  case KEY_ESC:
    exit(0);
    break;
  case 'f':
    if (fullscreen) {
      glutReshapeWindow(DEF_SCREEN_WIDTH, DEF_SCREEN_HEIGHT);
      glutPositionWindow(WINDOW_POS_X, WINDOW_POS_Y);
      fullscreen = 0;
    } else {
      glutFullScreen();
      fullscreen = 1;
    }
    break;
  }
}

void
idle()
{
  glutPostRedisplay();
}

static void
reshape(int w, int h)
{
  width = w;
  height = h;

  if (buffer) {
    free(buffer);
  }
  buffer = (GLubyte *)malloc(width * height * 4);
  if (!buffer) {
    abort();
  }
}

static void
draw_pixel(int x, int y, GLubyte r, GLubyte g, GLubyte b)
{
  int base = 4 * ((height - 1 - y) * width + x);
  buffer[base + 0] = r;
  buffer[base + 1] = g;
  buffer[base + 2] = b;
  buffer[base + 3] = 255;
}

static int block_size = 289;
static int symbol_hero[289] = 
                              {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,0,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,0,1,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,-1, 
                               0,0,0,0,0,0,0,1,1,0,0,1,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1}; 

static int symbol_skeleton[289] = 
                              {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,0,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,1,0,1,1,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,0,1,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,1,1,1,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,1,0,0,1,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1}; 

static int symbol_bat[289] = 
                              {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,0,1,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,1,1,1,1,1,1,1,1,0,0,0,-1, 
                               0,0,0,0,0,1,0,1,1,0,1,0,1,0,0,0,-1, 
                               0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1, 
                               0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1}; 

draw_symbol(int x, int y, int symbol[]) {
  GLubyte r,g,b;
  r = 255;
  int i,ri,ci,ti; 
  ci = 0;
  ri = 0;
  ti = block_size; 
  for(i = 0; i < ti; i++) {
    if(symbol[i] == -1) {
      ri++;
      ci = 0;
      continue;
    } else if(symbol[i] == 1) {
      draw_pixel(x + ci, y + ri, r,g,b);
      ci++;
    } else {
      ci++;
    }
  } 
}

static void
display()
{
  glClear(GL_COLOR_BUFFER_BIT);

  draw_symbol(100,100, symbol_hero);
  draw_symbol(150,150, symbol_skeleton);
  draw_symbol(120,120, symbol_skeleton);
  draw_symbol(130,100, symbol_bat);

  glWindowPos2i(0, 0);
  glDrawPixels(width, height, GL_RGBA, GL_UNSIGNED_BYTE, buffer);
  glutSwapBuffers();
}

int
main(int argc, char** argv)
{
  srand(0);
  reshape(DEF_SCREEN_WIDTH, DEF_SCREEN_HEIGHT);
  glutInit(&argc, argv);
  glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB | GLUT_DEPTH );
  glutInitWindowSize(DEF_SCREEN_WIDTH, DEF_SCREEN_HEIGHT);
  glutInitWindowPosition(WINDOW_POS_X, WINDOW_POS_Y);
  glutCreateWindow(argv[0]);
  glutKeyboardFunc(keyboard);
  glutIdleFunc(idle);
  glutReshapeFunc(reshape);
  glutDisplayFunc(display);

  glClearColor(0.0, 0.0, 0.0, 0.0);
  glShadeModel(GL_FLAT);

  glutMainLoop();
  return 0;
}
