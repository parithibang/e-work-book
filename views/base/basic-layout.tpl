{{.Header}}
<body>
	<!-- WRAPPER -->
	<div id="wrapper">
		<!-- NAVBAR -->
        {{template "base/sidebar.tpl" .}}
        <!-- MAIN -->
		<div class="main">
			<!-- MAIN CONTENT -->
			<div class="main-content">
				<div class="container-fluid">
					<!-- OVERVIEW -->
                    {{.LayoutContent}}
                    <!-- END OVERVIEW -->
				</div>
			</div>
			<!-- END MAIN CONTENT -->
		</div>
		<!-- END MAIN -->
{{.Footer}}