(module
    (func $print (import "myImports" "printFunc") (param i32))
    (func $add (param $a i32) (param $b i32) (result i32)
         local.get $a
         local.get $b
         i32.add)
    (func $main
        i32.const 40
        i32.const 2
        call $add
        call $print
    )
    (start $main)
)