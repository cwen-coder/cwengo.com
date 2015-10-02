(function() {
	var dom = {
		topic_del : $(".topic_del")
	}

	var del = {
		init : function() {
			this.eventFn();
		},
		eventFn : function() {
			dom.topic_del.on('click',function() {
				var url = "/admin/topic/DelTopic";
				var topic = $(this);
				var topic_id = topic.attr("section");
				$.post(url,{
					topic_id : topic_id
				},function(data) {
					alert(data['msg']);
					if(data['status'] == 1) {
						topic.parent().parent().remove();
					} 
				});
			})
		}
	}
	del.init();
})();

$(document).ready(function() {
	var liActive = $("div .pagination ul span").parent("li")
	liActive.addClass('active');
});