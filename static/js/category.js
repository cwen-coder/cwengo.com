(function() {
	var dom = {
		category_del : $(".category_del")
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

	del.init();
})();