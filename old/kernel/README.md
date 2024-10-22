# vmlinux.h

This file is needed for ebpf so "it" knows what exist in your kernel version. It can format 'raw'
and 'c', but we need 'go'

`bpftool btf dump file /sys/kernel/btf/vmlinux format c` outputs a lot, for instance:

The file `vmlinux` itself is in BTF format, which might be easier to parse, i.e. let bpftool output
Go code. OTOH, it doesn't look to bad.

also does JSON:

`bpftool btf dump -j file /sys/kernel/btf/vmlinux |jq`

~~~
[101] STRUCT 'file_system_type' size=72 vlen=17
        'name' type_id=5 bits_offset=0
        'fs_flags' type_id=21 bits_offset=64
        'init_fs_context' type_id=1181 bits_offset=128
        'parameters' type_id=1183 bits_offset=192
        'mount' type_id=1185 bits_offset=256
        'kill_sb' type_id=1158 bits_offset=320
        'owner' type_id=133 bits_offset=384
        'next' type_id=1008 bits_offset=448
        'fs_supers' type_id=85 bits_offset=512
        's_lock_key' type_id=122 bits_offset=576
        's_umount_key' type_id=122 bits_offset=576
        's_vfs_rename_key' type_id=122 bits_offset=576
        's_writers_key' type_id=1186 bits_offset=576
        'i_lock_key' type_id=122 bits_offset=576
        'i_mutex_key' type_id=122 bits_offset=576
        'invalidate_lock_key' type_id=122 bits_offset=576
        'i_mutex_dir_key' type_id=122 bits_offset=576
~~~

This is the first `struct`, in 'format c' this is generated:

~~~ c
struct file_system_type {
        const char *name;
        int fs_flags;
        int (*init_fs_context)(struct fs_context *);
        const struct fs_parameter_spec *parameters;
        struct dentry * (*mount)(struct file_system_type *, int, const char *, void *);
        void (*kill_sb)(struct super_block *);
        struct module *owner;
        struct file_system_type *next;
        struct hlist_head fs_supers;
        struct lock_class_key s_lock_key;
        struct lock_class_key s_umount_key;
        struct lock_class_key s_vfs_rename_key;
        struct lock_class_key s_writers_key[3];
        struct lock_class_key i_lock_key;
        struct lock_class_key i_mutex_key;
        struct lock_class_key invalidate_lock_key;
        struct lock_class_key i_mutex_dir_key;
};
~~~

Those `type_id` point to other types that we just need to use here. Not sure where `bits_offset`
comes into play, as the C code doesn't (explicitly) has it?

So in Go we want this to look as:

~~~ go
type FileSystemType struct {
        Name *string
        FsFlags int
        ...
~~~

I think I want native types as much as possible and only in the compile step make the size
assumptions.

## bpftool raw output

~~~
[101] STRUCT 'file_system_type' size=72 vlen=17
        'name' type_id=5 bits_offset=0
~~~

so
~~~
[ID] TYPE 'name' key=value key1=value1 ...
        'name' key2=value2 key3=value3
~~~

where some keys have more importance then others.

https://mostlynerdless.de/blog/2024/07/02/hello-ebpf-bpf-type-format-and-13-thousand-generated-java-classes-11/
