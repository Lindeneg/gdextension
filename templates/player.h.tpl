#ifndef {{.ProjectNameUpper}}_NODES_PLAYER_H_
#define {{.ProjectNameUpper}}_NODES_PLAYER_H_

#include <godot_cpp/classes/character_body2d.hpp>

#include "../core/macros.h"
{{if .WithLogger}}
#include "../objects/logger.h"
{{end}}

namespace godot {
class AnimatedSprite2D;
class CollisionShape2D;
}  // namespace godot

namespace godot::{{.ProjectName}} {

class Player : public CharacterBody2D {
    GDCLASS_EX(Player, CharacterBody2D)

    MDV_GET_SET(health, int, 100)
    MDV_GET_SET_T(sprite, AnimatedSprite2D*, nullptr);
    MDV_GET_SET_T(collision_shape, CollisionShape2D*, nullptr);

   private:
   {{if .WithLogger}}
    Ref<Logger> log_;
    {{end}}

   public:
    void _ready() override;
};
}  // namespace godot::{{.ProjectName}}

#endif  // {{.ProjectNameUpper}}_NODES_PLAYER_H_
