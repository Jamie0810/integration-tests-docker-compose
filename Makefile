run-migration:
	cd migration; \
  goose mysql "root:test@(127.0.0.1:3309)/kbc?parseTime=true" up;