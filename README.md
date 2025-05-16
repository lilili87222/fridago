# fridago
any platform 
https://github.com/aadog/frida_shared
https://github.com/aadog/frida-go


//要加参数 frida-go 编译的时候要加参数  -tags=tempdll,否则提示dll找不到
dll默认位置，请将binary_lib下的文件复制到下面位置
C:\Users\admin\AppData\Local\Temp/libfrida/24d920cc/frida_shared.dll
C:\Users\admin\AppData\Local\Microsoft\Windows\INetCache\frida\gadget-ios.dylib

不加-tags=tempdll则需要把frida_shared.dll复制到main.go同目录