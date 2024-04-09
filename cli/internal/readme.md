# `./internal`

The `internal` directory is primarily used to house packages that are internal to the vexal project.

Visibility Restriction: Packages within the internal directory are only accessible to other packages within the same Go module. They cannot be imported by packages outside of the module. This ensures that internal packages remain private to your project and are not used by external consumers.

Project-Specific Packages: The internal directory is suitable for packages that are tightly coupled with your project and not intended for reuse outside of it. These packages often contain implementation details or business logic specific to your project.

Enforcing Boundaries: By placing internal packages in the internal directory, you clearly delineate the boundaries between internal and external code. This makes it easier for developers to understand which parts of the codebase are meant for public consumption and which are not.
