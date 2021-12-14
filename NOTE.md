Naming convention 

### File
    1. All lower case with underscore separating multiple words
    2. File names that begin with “.” or “_” are ignored by the go tool 
    3. Files with the suffix _test.go are only compiled and run by the go test tool.
### function: 
    Public: PascalCase
    Private: camelCase
### interface: 
    MethodName + er = InterfaceName
### Unique name
    userID instead of userId
    productAPI instead of productApi
### variables
    Generally, use relatively simple (short) name.
    Consistent naming style should be used the entire source code
        - user to u
        - userID to uid
    If variable type is bool, its name should start with Has, Is, Can or Allow, etc.
    Single letter represents index: i, j, k

##Timeline
    Array, Slice
    Map, String
    Pointer
    Error
    Gorountine: dùng tạo concurrency -> go func()
    Channel: giúp giao tiếp giữa các gorountine
        + mutex (Locked(), Unlocked())
        + waitGroup(Add, Done, Wait)
        + Chanel sync