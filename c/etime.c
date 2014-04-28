#include <stdlib.h>
#include <stdio.h>
#include <sys/time.h>

void
tprint(char* prefix, long start, long stop)
{
double		usecs = (double) (stop - start);
double		secs = usecs / (double)1000000.0;

	fprintf(stderr, "%s%.4lf seconds\n", prefix, secs);
}

long
usecs()
{
	struct timeval tv;
	gettimeofday(&tv, NULL);
	return ( (long) tv.tv_sec * 1000000) + tv.tv_usec;
}

// Subtract the `struct timeval' values X and Y, toring the result in RESULT.
// Return 1 if the difference is negative, otherwise 0.
int
timeval_subtract(struct timeval *result, struct timeval *x, struct timeval *y) {
	// Perform the carry for the later subtraction by updating y.
	if (x->tv_usec < y->tv_usec) {
		int nsec = (y->tv_usec - x->tv_usec) / 1000000 + 1;
	 	y->tv_usec -= 1000000 * nsec;
	 	y->tv_sec += nsec;
	}
    if (x->tv_usec - y->tv_usec > 1000000) {
         int nsec = (x->tv_usec - y->tv_usec) / 1000000;
         y->tv_usec += 1000000 * nsec;
         y->tv_sec -= nsec;
    }
     
	// Compute the time remaining to wait. tv_usec is certainly positive.
	result->tv_sec = x->tv_sec - y->tv_sec;
	result->tv_usec = x->tv_usec - y->tv_usec;
     
    // Return 1 if result is negative.
	return x->tv_sec < y->tv_sec;
}
