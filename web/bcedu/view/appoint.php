<!DOCTYPE html>
<html>
	<head>
		<title>登记预约</title>
		<meta name=\"description\" content=\"登记预约。填写预约订单，写清楚您的需求，就可以获得免费匹配老师，试体验一次服务。\">
		<meta name=\"keywords\" content=\"登记预约,免中介信息费,去中心化，量子教育网\">
<?php
include(V.'header-content.php');
?>
	</head>
	<body>
<?php
//导航条
include(V.'top-bar.php');
?>
<?php
		$GRADES=array(
				"小一",
				"小学一年级",
				"研三",
				"博士",
				"其它年级"
				);
?>


		<div class="container">
			<div class="col-sm-6 col-sm-offset-3">
				<?php
				$GRADES=array(
				"一年级",
				"二年级",
				"三年级",
				"四年级",
				"五年级",
				"六年级",
				"初一",
				"初二",
				"初三",
				"高一",
				"高二",
				"高三",
				"其它年级"
				);
				$DISTRICTS=array(
				"包河区",
				"瑶海区",
				"蜀山区",
				"高新区",
				"其他",
				);

				$SUBJECTS=array(
				"小学语文",
				"小学数学",
				"小学英语",
				"初中语文",
				"初中数学",
				"初中英语",
				"高中语文",
				"高中数学",
				"高中英语",
				"初中物理",
				"初中化学",
				"高中物理",
				"高中化学",
				"史地政生",
				"竞赛",
				"复读",
				"课外英语",
				"艺术类",									  
				"其它"
				);
				?>
				<form role="form" action="?page=appointed" method="post">
					<h3>学生登记：</h3>
					<div class="panel panel-info">

						<div class="panel-heading">
							<h3 class="panel-title">#</h3>
						</div>
						<div class="panel-body">
							  <table class="table" style="table-layout:fixed">
							  <h4>填写预约信息</h4>
								  <tr>
									<td><label>姓名：*</label><input type="text" name="name" class="form-control"></td>
									<td>
										<label>年级：</label>
											<select name="grade" class="form-control">
											<?php 
												foreach ($GRADES as $GRADE)
												{
													echo "<option value='$GRADE'>$GRADE</option>";
												}
											?>
											</select>
									</td>
									<td>
										<label>性别：</label>
										<select name="gender" class="form-control">
												<option value="男生">男生</option>
												<option value="女生">女生</option>
											</select>
									</td>
								  </tr>
								  <tr>
									<td>
										<label>城市：</label>
											<select name="address" class="form-control" disabled>
											<option selected='selected'>合肥</option>";
											</select>
									</td>
									<td>
										<label>住址：</label>
											<select name="address" class="form-control">
											<?php 
												foreach ($DISTRICTS as $DISTRICT)
												{
													echo "<option value='$DISTRICT'>$DISTRICT</option>";
												}
											?>
											</select>
									</td>
								  </tr>
								  <tr>
								  <td colspan="3">
									<label>详细住址（街道/小区）：</label>
									<textarea name="detailed" placeholder="" 
										class="form-control"></textarea>
									</td>
								  </tr>
								  <tr>
									<td colspan="2"><label>联系电话：*</label><input type="number" name="tel" class="form-control"></td>
								  </tr>
								  <tr>
									<td colspan="3">
										<label>补习科目：</label>
										<?php
										foreach ($SUBJECTS as $SUBJECT)
											{
												echo "<label class='checkbox-inline'>
													<input type='checkbox' name='subject[]' value='$SUBJECT'> $SUBJECT
													</label>";
											}
										?>
									</td>
								  </tr>

									<tr>
									<td colspan="3">
									<label>时间及薪酬：</label>
									<textarea name="timepay" placeholder="每小时薪酬×元" 
										class="form-control"></textarea>
									</td>
								  </tr>

								  <tr>
									<td colspan="3">
									<label>对家教的要求：</label>
									<textarea name="want" placeholder="例如：具体需要教授内容和学习程度等" 
										class="form-control"></textarea>
									</td>
								  </tr>
								  <tr>

								  </tr>
							  </table>
						</div>
					</div>
					<div style="text-align:center">
						<a href="?page=student"><button type="button" class="btn btn-info" style="float:left">返回</button></a>
						<button type="submit" class="btn btn-primary">确认提交</button>
					</div>
				</form>

<?php
//页尾
include(V.'footer-content.php');
?>
	</body>
</html>