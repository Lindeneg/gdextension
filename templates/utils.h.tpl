#ifndef {{.ProjectNameUpper}}_UTILS_H_
#define {{.ProjectNameUpper}}_UTILS_H_

#include <godot_cpp/classes/node.hpp>

namespace godot {
class String;
class StringName;
class Callable;
class Resource;
}  // namespace godot

namespace godot::{{.ProjectName}}::utils {

template <typename T>
T *create_component(Node *owner, bool add_child = true) {
    T *obj = memnew(T);
    if (add_child) {
        owner->add_child(obj);
        obj->set_owner(owner);
    }
    return obj;
}

bool is_in_editor();
bool is_in_game();
const char *convert_gd_string(const Node *n);
const char *convert_gd_string(const Resource *n);
const char *convert_gd_string(String s);
const char *convert_gd_string(StringName s);
void connect(Node *node, const StringName name, const Callable &callable);
void disconnect(Node *node, const StringName name, const Callable &callable);
void queue_delete(Node *node);
bool is_pressed(StringName name);
}  // namespace godot::{{.ProjectName}}::utils

#endif  // {{.ProjectNameUpper}}_UTILS_H_

