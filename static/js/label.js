(function() {
	var dom = {
		label_del : $(".label_del"),
		label_edit : $(".label_edit")
 	}

	var del = {
		init : function() {
			this.eventFn();
		},

		eventFn : function() {
			dom.label_del.on('click',function() {
				var url = "/admin/label/DelLabel";
				var label = $(this);
				var label_id = label.attr("section");
				$.post(url,{
					label_id : label_id
				},function(data) {
					alert(data['msg']);
					if(data['status'] == 1) {
						label.parent().parent().remove();
					} 
				});
			})
		}
	}

	var edit = {
		init : function() {
			this.eventFn();
		},

		eventFn : function() {
			dom.label_edit.on('click',function() {
				var label = $(this).parent().parent().children();
				var label_id = label.eq(0).text();
				var label_name = label.eq(1).text();
				var editLabel = $('#editLabel');
				editLabel.modal('show');
				editLabel.find('.modal-body #EditNewLabel').val(label_name);
				editLabel.find('.modal-footer .btn-primary').on('click',function() {
					var url = "/admin/label/EditLabel";
					$.post(url,{
						label_id : label_id,
						label_name : editLabel.find('.modal-body #EditNewLabel').val()
					},function(data) {
						alert(data['msg']);
						if(data['status'] == 1) {
							window.location.reload();
						}
					
					});
				});
			});
		}
	}


	del.init();
	edit.init();
})();