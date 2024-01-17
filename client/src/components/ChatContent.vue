<template>
  <div id="container" class="dialogue" type="flex" @scroll="scroll">
		<div v-for="message, i in messageList" :key="i">
			<div class="date">
				{{ dateList[i] }}
			</div>
			<div v-if="message.name == userName" class="user local">
				<div class="msg">
					<div class="text">{{ message.text }}</div>
					<div class="time">{{ convertToTime(message.time) }}</div>
				</div>
				<div class="avatar">
					<div class="pic">
						<img src="../assets/avatar1.jpg">
					</div>
					<div class="name">{{ message.name }}</div>
				</div>
			</div>
			<div v-else class="user remote">
				<div class="avatar">
					<div class="pic">
						<img src="../assets/avatar2.jpg">
					</div>
					<div class="name">{{ message.name }}</div>
				</div>
				<div class="msg">
					<div class="text">{{ message.text }}</div>
					<div class="time">{{ convertToTime(message.time) }}</div>
				</div>
			</div>
		</div>
  </div>
</template>

<script>
export default {
  name: 'ChatContent',
  props: {
		userName: String,
    messageList: Array,
  },

	data() {
		return {
			dateList: [],
			scrollTop: 0,
			maxScrollTop: 0
		}
	},

	mounted() {
		this.autoScroll();
	},

	methods: {
		scroll(sender) {
			this.scrollTop = sender.target.scrollTop;
			if (this.scrollTop == 0) {
				this.getMoreMessage();
			}
		},

		autoScroll() {
			var dialogue = document.querySelector("#container")
			var maxScrollTop = dialogue.scrollHeight - dialogue.clientHeight;
			if (this.scrollTop == this.maxScrollTop) {
				this.$nextTick(() => {
					dialogue.scrollTop = dialogue.scrollHeight
				});
			} else {
				if (this.scrollTop == 0) {
					this.$nextTick(() => {
						dialogue.scrollTop = maxScrollTop - this.maxScrollTop
					});
				}
			}
			this.$nextTick(() => {
				this.maxScrollTop = dialogue.scrollHeight - dialogue.clientHeight;
			});
		},

		getMoreMessage() {
			this.$emit("history");
		},

		newDateList() {
			this.dateList = [];
			if (this.messageList != null) {
				this.messageList.forEach((message) => {
					var date = this.convertToDate(message.time);
					if (this.dateList.indexOf(date) == -1) {
						this.dateList.push(date);
					} else {
						this.dateList.push("");
					}
				});
			}
		},

		convertToDate(time) {
			var messageTime = new Date(time);
			var date = messageTime.getDate();
			var month = messageTime.getMonth() + 1;
			var year = messageTime.getFullYear();
			return `${year}/${month}/${date}`;
		},

		convertToTime(time) {
			var date = new Date(time);
			var hour = date.getHours();
			if (hour < 10) {
				hour = "0" + hour;
			}
			var minute = date.getMinutes();
			if (minute < 10) {
				minute = "0" + minute;
			}

			return hour + ":" + minute;
		}
	},

	watch: {
		messageList: {
			handler() {
				this.autoScroll();
				this.newDateList();
			}
		}
	}
}
</script>

<style lang="scss">
* {
	margin: 0;
	padding: 0;
	list-style: none;
}

html,
body {
	height: 100%;
}

.dialogue {
  height: 100%;
	box-shadow: 0 0 1px #000;
	background-color: #f4f5f7;
	overflow-x: hidden;
	overflow-y: scroll;
}

.date {
	font-size: 25px;
	font-weight: bolder;
	text-align: center;
	margin-top: 10px;
	margin-bottom: 10px;
}

.user {
	bottom: 100%;
	display: flex;
	align-items: flex-start;
	.avatar {
		width: 50px;
		text-align: center;
		flex-shrink: 0;
	}
	.pic {
		border-radius: 50%;
		overflow: hidden;
		img {
			width: 100%;
			vertical-align: middle;
		}
	}
	.name {
		color: #333;
	}
	.msg {
		margin-top: 2px;
		position: relative;
	}
	.text {
		background-color: #aaa;
		padding: 15px;
		border-radius: 10px;
		position: relative;
	}
	.time {
		font-size: 13px;
		margin-top: 2px;
		color: #aaa;
	}
}

.remote {
	.msg {
		margin-left: 20px;
		margin-right: 80px;
	}
	.text {
		color: #eee;
		background-color: #4179f1;
		&::before {
			border-right: 10px solid #4179f1;
			left: -10px;
		}
	}
	.time {
		margin-left: 5px;
		text-align: left;
	}
}
.local {
	justify-content: flex-end;
	.msg {
		margin-right: 20px;
		margin-left: 80px;
	}
	.text {
		order: -1;
		background-color: #fff;
		color: #333;
		&::before {
			border-left: 10px solid #fff;
			right: -10px;
		}
	}
	.time {
		margin-right: 5px;
		text-align: right;
	}
}

.remote,
.local {
	& .text::before {
		content: "";
		position: absolute;
		top: 20px;
		border-top: 10px solid transparent;
		border-bottom: 10px solid transparent;
	}
	.text {
		font-weight: lighter;
		box-shadow: 0 0 10px #888;
	}
}
</style>
