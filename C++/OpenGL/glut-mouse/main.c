#include <GL/glut.h>

int window_w, window_h = 0;
float mouse_x, mouse_y = 0.0;
float box_size = 0.02;

void display(void)
{
  glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);

  glBegin(GL_POLYGON);
  glVertex3f(mouse_x+box_size, mouse_y+box_size, 0.0);
  glVertex3f(mouse_x-box_size, mouse_y+box_size, 0.0);
  glVertex3f(mouse_x-box_size, mouse_y-box_size, 0.0);
  glVertex3f(mouse_x+box_size, mouse_y-box_size, 0.0);
  glEnd();

  glutSwapBuffers();
}

void motion(int x, int y)
{
  mouse_x = (float)x/window_w - 0.5;
  mouse_y = -(float)y/window_h + 0.5;
  glutPostRedisplay();
}

void reshape(int w, int h)
{
  window_w = w;
  window_h = h;
  glViewport(0, 0, w, h);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  gluOrtho2D(-.5, .5, -.5, .5);
}

int main(int argc, char **argv)
{
  glutInit(&argc, argv);
  glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB | GLUT_DEPTH);
  glutInitWindowSize(500, 500);
  glutCreateWindow("Square Follows Mouse - GLUT");
  glutPassiveMotionFunc(motion);
  glutReshapeFunc(reshape);
  glutDisplayFunc(display);
  glutMainLoop();
  return 0;
}
