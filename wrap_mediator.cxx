#include "wrap_mediator.hxx"
#include "mediator.hxx"

int add_from_cplus(int a, int b) {
  Summer num;
  return num.add(a, b);
}
