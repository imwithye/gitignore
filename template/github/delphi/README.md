Git Template Delphi
===

Use `git ignore add delphi` to add this ignore template.

```
# Uncomment these types if you want even more clean repository. But be careful.
# It can make harm to an existing project source. Read explanations below.
#
# Resource files are binaries containing manifest, project icon and version info.
# They can not be viewed as text or compared by diff-tools. Consider replacing them with .rc files.
#*.res
#
# Type library file (binary). In old Delphi versions it should be stored.
# Since Delphi 2009 it is produced from .ridl file and can safely be ignored.
#*.tlb
#
# Diagram Portfolio file. Used by the diagram editor up to Delphi 7.
# Uncomment this if you are not using diagrams or use newer Delphi version.
#*.ddp
#
# Visual LiveBindings file. Added in Delphi XE2.
# Uncomment this if you are not using LiveBindings Designer.
#*.vlb
#
# Deployment Manager configuration file for your project. Added in Delphi XE2.
# Uncomment this if it is not mobile development and you do not use remote debug feature.
#*.deployproj
# 
# C++ object files produced when C/C++ Output file generation is configured.
# Uncomment this if you are not using external objects (zlib library for example).
#*.obj
#

# Delphi compiler-generated binaries (safe to delete)
*.exe
*.dll
*.bpl
*.bpi
*.dcp
*.so
*.apk
*.drc
*.map
*.dres
*.rsm
*.tds
*.dcu
*.lib
*.a
*.o
*.ocx

# Delphi autogenerated files (duplicated info)
*.cfg
*.hpp
*Resource.rc

# Delphi local files (user-specific info)
*.local
*.identcache
*.projdata
*.tvsconfig
*.dsk

# Delphi history and backups
__history/
__recovery/
*.~*

# Castalia statistics file (since XE7 Castalia is distributed with Delphi)
*.stat
```