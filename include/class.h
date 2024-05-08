#include "omp.h"

#ifdef __cplusplus
extern "C"
{
#endif

	typedef struct ClassData
	{
		int team;
		int skin;
		float spawnX;
		float spawnY;
		float spawnZ;
		float angle;
		uint8_t weapon1;
		uint32_t ammo1;
		uint8_t weapon2;
		uint32_t ammo2;
		uint8_t weapon3;
		uint32_t ammo3;
	} ClassData;

	void* class_create(ClassData* data);
	void class_release(void* _class);
	int class_getID(void* class_);
	void class_setClass(void* class_, ClassData* data);
	ClassData class_getClass(void* class_);

#ifdef __cplusplus
}
#endif
