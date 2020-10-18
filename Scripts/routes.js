var func = require('./function.js');

module.exports = function(app) {
	app.post('/cr_user', function(req, res) {
		func.RegisterNewUser(req, res);
	});
	app.post('/cr_cal', function(req, res) {
		func.AddNewMR(req, res);
	});
	app.post('/se_cal', function(req, res) {
		func.GetMRByID(req, res);
	});
	app.get('/cr_admin', function(req, res) {
		func.RegisterNewAdmin(req, res);
	});
}
