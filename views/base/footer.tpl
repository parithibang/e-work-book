<div class="clearfix"></div>
		<footer>
			<div class="container-fluid">
				<p class="copyright">&copy; 2020 <a href="#" target="_blank">GIBS</a>. All Rights Reserved.</p>
			</div>
		</footer>
	</div>
	<!-- END WRAPPER -->
	<!-- Javascript -->
	<script src="/static/vendor/jquery/jquery.min.js"></script>
	<script src="/static/vendor/bootstrap/js/bootstrap.min.js"></script>
	<script src="/static/vendor/jquery-slimscroll/jquery.slimscroll.min.js"></script>
	<script src="/static/vendor/jquery.easy-pie-chart/jquery.easypiechart.min.js"></script>
	<script src="/static/vendor/chartist/js/chartist.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.9/dist/js/bootstrap-select.min.js"></script>
	<script src="/static/js/klorofil-common.js"></script>
    <script src="/static/js/common.js"></script>
    <script src="/static/js/common.js"></script>
     <script>
	$(function() {	
		var data, options;	
		// headline charts	
		data = {	
			labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],	
			series: [	
				[23, 29, 24, 40, 25, 24, 35],	
				[14, 25, 18, 34, 29, 38, 44],	
			]	
		};	
		options = {	
			height: 300,	
			showArea: true,	
			showLine: false,	
			showPoint: false,	
			fullWidth: true,	
			axisX: {	
				showGrid: false	
			},	
			lineSmooth: false,	
		};
        if ($("#headline-chart").length){
            new Chartist.Line('#headline-chart', data, options);
        }		
	});	
	</script>
</body>

</html>
