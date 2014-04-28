// Copyright © 2014 Lawrence E. Bakst. All rights reserved.

// simple typedefs to get bool, Int, Str
typedef char bool;
enum bool {false, true};
typedef long int Int;
typedef char* Str;

// simple Go style interface for sort
typedef struct Interface Interface;
struct Interface {
	Int (*Len)(Interface *d);
	bool (*Less)(Interface *d, Int i, Int j);
	void (*Swap)(Interface *d, Int i, Int j);
	bool (*Add)(Interface *d, void* p, Int len);
	void* (*Get)(Interface *d, Int i);
	bool (*Set)(Interface *d, Int i, void* p);
	Int (*MaxLen)(Interface *d);
	bool (*Free)(Interface *d);
	void* (*data);
	Int dataLen;		// if 0 use strcmp or int compare, nothing to free, if nonzero use strncmp and Free on Del.
	Int maxLen;			// maximum lenbth of data array
	Int len;			// current length of data array
};

// External API
void		HeapSort(Interface *d);

// Debug API
void		Pl(Interface *d, char* str);

// Generic functions for Interface
Int			Len(Interface *d);
void		Swap(Interface *d, Int i, Int j);
void*		Get(Interface *d, Int i);
Int			MaxLen(Interface *d);

//Generic API for Interface
void		DelInterface(Interface* d);

// External API for type Int
Interface*	NewIntInterface(Int n);
bool		IntLess(Interface *d, Int i, Int j);
bool		IntAdd(Interface *d, void* p, Int i);

// External API for type Str
Interface *	NewStrInterface(Int n, Int dataLen);
bool		StrLess(Interface *d, Int i, Int j);
bool 		StrAdd(Interface *d, void* p, Int len);
bool		StrFree(Interface *d);
Str			NewStr(Int len);

// internal functions
void		siftup(Interface *d, Int ni, Int ri);
