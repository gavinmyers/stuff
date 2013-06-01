#include <stdio.h>
#include <stdlib.h>
//#include <GLUT/glut.h> 
#include <GL/glut.h>//Drawing funciton

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

static void
display()
{
  int x, y, xi, yi;
  int sq = 10;
  int ri = 255;

  glClear(GL_COLOR_BUFFER_BIT);

  GLubyte r =  255;
  GLubyte g =  255;
  GLubyte b =  255;
  
  for (y = 0; y < height; y = y + sq) {
    for (x = 0; x < width; x = x + sq) {
      r = rand() % 255;
      g = rand() % 10; 
      b = ri;
      for(yi = 0; yi < sq; yi++) {
        for(xi = 0; xi < sq; xi++) {
          draw_pixel(x + xi, y + yi, r, g, b);
        }
      }
      if(ri == 255) {
        ri = 0;
      } else {
        ri = 255;
      }
    }
  }

  glWindowPos2i(0, 0);
  glDrawPixels(width, height, GL_RGBA, GL_UNSIGNED_BYTE, buffer);
  glutSwapBuffers();
}

int
main(int argc, char** argv)
{
  srand(0);

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
