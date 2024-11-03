#include "utils.h"

#include <godot_cpp/classes/engine.hpp>
#include <godot_cpp/classes/global_constants.hpp>
#include <godot_cpp/classes/input.hpp>
#include <godot_cpp/classes/object.hpp>
#include <godot_cpp/classes/resource.hpp>
#include <godot_cpp/core/class_db.hpp>
#include <godot_cpp/core/error_macros.hpp>
#include <godot_cpp/variant/callable.hpp>

namespace godot::{{.ProjectName}} {
bool utils::is_in_editor() {
    return godot::Engine::get_singleton()->is_editor_hint();
}

bool utils::is_in_game() { return !(is_in_editor()); }

const char *utils::convert_gd_string(const Node *n) {
    ERR_FAIL_NULL_V_MSG(n, "", "cannot get_name on node that is nullptr");
    return convert_gd_string(n->get_name());
}

const char *utils::convert_gd_string(const Resource *n) {
    ERR_FAIL_NULL_V_MSG(n, "", "cannot get_name on resource that is nullptr");
    return convert_gd_string(n->get_name());
}

const char *utils::convert_gd_string(StringName s) {
    return convert_gd_string(String(s));
}

const char *utils::convert_gd_string(String s) { return s.utf8().get_data(); }

void utils::connect(Node *node, const StringName name,
                    const Callable &callable) {
    ERR_FAIL_NULL_MSG(node, vformat("cannot connect to signal %s", name));
    if (!node->is_connected(name, callable)) {
        if (node->connect(name, callable) != OK) {
            WARN_PRINT_ED(vformat("%s failed to connect to signal %s",
                                  node->get_name(), name));
        }
    }
}

void utils::disconnect(Node *node, const StringName name,
                       const Callable &callable) {
    if (!node) return;
    if (node->is_connected(name, callable)) {
        node->disconnect(name, callable);
    }
}

void utils::queue_delete(Node *node) {
    if (!node || node->is_queued_for_deletion()) {
        return;
    }
    node->queue_free();
}

bool utils::is_pressed(StringName name) {
    return Input::get_singleton()->is_action_pressed(name);
}
}  // namespace godot::{{.ProjectName}}


