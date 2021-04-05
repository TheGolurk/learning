#include <stdexcept>
#include <limits>
#include <iostream>

using namespace std;

void Compare(int c) {
	if (c > numeric_limits< char>::max())
		throw invalid_argument("Compare argument too large");
}

int main () {
		
	try {
		Compare(256);
	} catch (invalid_argument& e) {
		cerr << e.what() << endl;
		return -1;
	}

	return 0;
}
