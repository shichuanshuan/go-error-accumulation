【】
来自 exec.Command

sysattr := &syscall.ProcAttr{
Dir: attr.Dir,
Env: attr.Env,
Sys: attr.Sys,
}
if sysattr.Env == nil {
sysattr.Env, err = execenv.Default(sysattr.Sys)
if err != nil {
return nil, err
}
}
sysattr.Files = make([]uintptr, 0, len(attr.Files))
for _, f := range attr.Files {
sysattr.Files = append(sysattr.Files, f.Fd())
}

为什么结构体不直接赋值，确实是用 append 呢

因为切片是引用类型，要是赋值，那么相当于是对同一地址操作