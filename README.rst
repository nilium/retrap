retrap
------
::

    retrap [SRC:DST] [SRC:-] -- CMD [ARG...]

retrap is a small program that spawns a subprocess and forwards remapped signals
to it.

Signals are mapped by passing `SRC:DST` flags before the `--` separator, where
*SRC* is the signal to trap in the parent process and *DST* is the signal to
send to the subprocess. So, `INT:TERM` will trap an interrupt signal to the
parent process and send a terminate signal to the subprocess (the *CMD*).

In addition, a signal can be discarded by specifying `SRC:-`, where *DST* is
a hyphen. For example, if `INT:-` were passed, the parent process would trap and
discard interrupt signals.

By default, all signals are sent to the subprocess unless mapped.

This is a weird, niche little tool, but can be useful for dealing with programs,
on the CLI, that expect specific signals. For example, if fiddling with runsv,
you can use `retrap int:term -- runsv service` to translate an interrupt into
a terminate, telling runsv to tear down (the original reason this was written).

retrap is not guaranteed to work perfectly in all cases, and probably can't. It
cannot trap the usual signals (SIGKILL and SIGSTOP on Linux, for example). It
will ignore SIGPIPE, simply forwarding that on to its subprocess.

The subprocess receives the parent process's standard input, output, and error
file descriptors.

.. vim: set ft=rst tw=80 sw=4 ts=4 et :
