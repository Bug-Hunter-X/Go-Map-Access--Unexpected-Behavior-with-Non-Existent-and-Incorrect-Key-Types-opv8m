# Go Map Access: Unexpected Behavior with Non-Existent and Incorrect Key Types

This example demonstrates unexpected behaviors when accessing Go maps with non-existent keys and incorrect key types.

Accessing a map with a key that does not exist will not trigger a panic; it will return the zero value of the map's value type (0 for `int`). Accessing with an incorrect key type will cause a panic.

The solution showcases a way to handle these situations by checking for the existence of the key before accessing the map's value.