# frida_shared.dll 生成教程

本库是一个将 frida-core 转换为 frida_shared.dll 的自动化脚本。

## 快速开始

1. 启动 `build_windows.bat`，自动生成 frida_shared.dll。
2. `build_windows.bat` 会调用 `pre_windows.py` 下载 frida_core 发布版本。
3. `pre_windows.py` 会调用 `builddef.py` 生成 `import.c` 和 `frida.def` 两个文件。
4. `builddef.py` 生成 `import.c` 时依赖 libclang，请确保 Python 的 clang 库和本地安装的 libclang.dll 版本一致。

   安装方式示例：
   ```shell
   python.exe -m pip install --upgrade pip
   pip install libclang==18.*  # 若本地 clang 为 18，则需同步安装 18 版本的 python 库
   ```

5. 参考原编译教程：https://github.com/aadog/frida_shared/blob/master/README.MD

## 编译前准备

1. 安装 LLVM 或 Clang，并修改 `builddef.py` 里的 libclang 路径。
2. 修改 `frida_ver.py` 中的版本号为最新（如遇问题，可手动编辑 __pycache__ 里的版本号）。
   - github 原项目版本为 15.1.21，本项目尝试更新到 frida 16.7。
3. 修改 `build_windows.bat` 里的 VS 安装路径。

## 常见问题

- **编译报错：找不到 SetupDiGetDeviceInstanceIdW**
  - 解决方法：在 `CMakeLists.txt` 中添加：
    ```cmake
    if(WIN32)
        target_link_libraries(frida_shared setupapi)
    endif()
    ```

> 原项目版本 ver="15.1.21"，本项目尝试更新到 frida 16.7。