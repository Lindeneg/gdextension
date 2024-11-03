#ifndef {{.ProjectNameUpper}}_CORE_MACROS_H_
#define {{.ProjectNameUpper}}_CORE_MACROS_H_

#include <godot_cpp/core/binder_common.hpp>
#include <godot_cpp/core/class_db.hpp>

#include "./utils.h"

#define GDSTR(s) utils::convert_gd_string(s)

#define MAKE_RESOURCE_TYPE_HINT(m_type) \
    vformat("%s/%s:%s", Variant::OBJECT, PROPERTY_HINT_RESOURCE_TYPE, m_type)

#define FIND_CHILD(name, pattern, type)                                      \
    {                                                                        \
        Node *child{find_child(pattern)};                                    \
        ERR_FAIL_NULL_MSG(child, "required component " #name " is missing"); \
        name = static_cast<type *>(child);                                   \
    }

#define FIND_OR_CREATE_CHILD(name, pattern, type)                  \
    {                                                              \
        if (!name) {                                               \
            type *child{static_cast<type *>(find_child(pattern))}; \
            if (child) {                                           \
                name = static_cast<type *>(child);                 \
            } else {                                               \
                name = utils::create_component<type>(this);        \
                name->set_name(pattern);                           \
            }                                                      \
        }                                                          \
    }

#define REQUIRED_CHILD(name, pattern, type)        \
    if (utils::is_in_editor()) {                   \
        FIND_OR_CREATE_CHILD(name, pattern, type); \
    } else {                                       \
        FIND_CHILD(name, pattern, type);           \
    }

// macros to ease the godots boilerplates
// and for exposing properties to editor

#define GDCLASS_EX(cls, base) \
    GDCLASS(cls, base)        \
   public:                    \
    cls();                    \
    ~cls();                   \
                              \
   protected:                 \
    static void _bind_methods();

// getter
#define M_GET(name, type) \
    inline type get_##name() const { return name##_; }

// const setter
#define M_SET(name, type) \
    inline void set_##name(const type n) { name##_ = n; }

// const ref setter
#define M_SET_R(name, type) \
    inline void set_##name(const type &n) { name##_ = n; }

// T setter
#define M_SET_T(name, type) \
    inline void set_##name(type n) { name##_ = n; }

// declares property
#define P_DECL(name, type) \
   private:                \
    type name##_;

// declares and initializes property
#define P_DECL_V(name, type, value) \
   private:                         \
    type name##_ = value;

// declaration and getter
#define MD_GET(name, type) \
    P_DECL(name, type)     \
   public:                 \
    M_GET(name, type)

// declaration, initialazation and getter
#define MDV_GET(name, type, value) \
    P_DECL_V(name, type, value)    \
   public:                         \
    M_GET(name, type)

// getter and const setter
#define M_GET_SET(name, type) \
   public:                    \
    M_GET(name, type)         \
    M_SET(name, type)

// declaration, getter and const setter
#define MD_GET_SET(name, type) \
    P_DECL(name, type)         \
    M_GET_SET(name, type)

// declaration, initialazation, getter and const setter
#define MDV_GET_SET(name, type, value) \
    P_DECL_V(name, type, value)        \
    M_GET_SET(name, type)

// getter and const ref setter methods
#define M_GET_SET_R(name, type) \
   public:                      \
    M_GET(name, type)           \
    M_SET_R(name, type)

// declaration, getter and const ref setter methods
#define MD_GET_SET_R(name, type) \
    P_DECL(name, type)           \
    M_GET_SET_R(name, type)

// declaration, initialazation, getter and const ref setter methods
#define MDV_GET_SET_R(name, type, value) \
    P_DECL_V(name, type, value)          \
    M_GET_SET_R(name, type)

// getter and T setter methods
#define M_GET_SET_T(name, type) \
   public:                      \
    M_GET(name, type)           \
    M_SET_T(name, type)

// declaration, getter and T setter methods
#define MD_GET_SET_T(name, type) \
    P_DECL(name, type)           \
    M_GET_SET_T(name, type)

// declaration, initialzation, getter and T setter methods
#define MDV_GET_SET_T(name, type, value) \
    P_DECL_V(name, type, value)          \
    M_GET_SET_T(name, type)

#define M_DEBUG() MDV_GET_SET(debug, bool, false)

#define M_BIND(name, cls)                                           \
    ClassDB::bind_method(D_METHOD("get_" #name), &cls::get_##name); \
    ClassDB::bind_method(D_METHOD("set_" #name, #name), &cls::set_##name);

#define MP_BIND(name, cls, prop) \
    M_BIND(name, cls)            \
    ClassDB::add_property(#cls, prop, "set_" #name, "get_" #name);

// bind variant
#define MPV_BIND(name, cls, prop) MP_BIND(name, cls, PropertyInfo(prop, #name))

// bind node
#define MPN_BIND(name, cls, prop) \
    MP_BIND(                      \
        name, cls,                \
        PropertyInfo(Variant::OBJECT, #name, PROPERTY_HINT_NODE_TYPE, #prop))

// bind resource
#define MPR_BIND(name, cls, prop) \
    MP_BIND(name, cls, PropertyInfo(Variant::OBJECT, #prop))

#define MP_DEBUG_BIND(cls) MPV_BIND(debug, cls, Variant::BOOL)

#endif  // {{.ProjectNameUpper}}_CORE_MACROS_H_

