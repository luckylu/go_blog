<!DOCTYPE html>
<html>
	<head>
		{{template "layout/head.tpl" .}}
		{{template "head" .}}
	</head>
	<body>
		{{template "layout/sidebar.tpl" .}}
		{{template "body" .}}
		{{template "layout/footer.tpl" .}}
	</body>
	<script>
		$('.sidebar').sidebar({
			context: '.pushable'
		})

		$('.dropdown').dropdown();
	</script>
	<style>
		.pushable {
			height: auto!important;
		}

		.pusher {
		  overflow: visible!important;
		}
	</style>
</html>
