#ifndef {{.ProjectNameUpper}}_EXTENSION_REGISTER_TYPES_H_
#define {{.ProjectNameUpper}}_EXTENSION_REGISTER_TYPES_H_

#include <godot_cpp/core/class_db.hpp>

using namespace godot;

void initialize_{{.ProjectName}}_extension_module(ModuleInitializationLevel p_level);
void uninitialize_{{.ProjectName}}_extension_module(ModuleInitializationLevel p_level);

#endif  // {{.ProjectNameUpper}}_EXTENSION_REGISTER_TYPES_H_

