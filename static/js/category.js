(function() {
	var dom = {
		category_del : $(".category_del"),
		category_edit : $(".category_edit")
 	}

	var del = {
		init : function() {
			this.eventFn();
		},

		eventFn : function() {
			dom.category_del.on('click',function() {
				var url = "/admin/category/DelCategory";
				var category = $(this);
				var category_id = category.attr("section");
				$.post(url,{
					category_id : category_id
				},function(data) {
					alert(data['msg']);
					if(data['status'] == 1) {
						category.parent().parent().remove();
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
			dom.category_edit.on('click',function() {
				var category = $(this).parent().parent().children();
				var category_id = category.eq(0).text();
				var category_name = category.eq(1).text();
				var editCategory = $('#editCategory');
				editCategory.modal('show');
				editCategory.find('.modal-body #EditNewCategory').val(category_name);
				editCategory.find('.modal-footer .btn-primary').on('click',function() {
					var url = "/admin/category/EditCategory";
					$.post(url,{
						category_id : category_id,
						category_name : editCategory.find('.modal-body #EditNewCategory').val()
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