# Emacs
## Toggle outline element to new window

```lisp
(defun split-and-indirectbuffer-orgtree ()
"Splits current window to the right and opens the selected section in it as an indirect buffer"
(interactive)
(split-window-right)
(org-tree-to-indirect-buffer)
(windmove-right))

(defun kill-and-unsplit-indirectbuffer-orgtree ()
"Kills the indirect buffer and deletes the window."
(interactive)
(kill-this-buffer)
(delete-window))

(with-eval-after-load 'org
	(define-key org-mode-map [f9] #'split-and-indirectbuffer-orgtree)
	(define-key org-mode-map [f12] #'kill-and-unsplit-indirectbuffer-orgtree))
```
