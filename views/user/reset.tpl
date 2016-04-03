<br/>
<div class="forgotmodal-dialog">
    <div class="forgotmodal-container">
        <div id="brand">&nbsp</div>
        &nbsp;
        <h3 class="dark-grey">Reset Password</h3>
        &nbsp;
        &nbsp;
        {{if .flash.error}}
        <h3>{{.flash.error}}</h3>
        &nbsp;
        {{end}}
        {{if .Errors}}
        {{range $rec := .Errors}}
        <h3>{{$rec}}</h3>
        {{end}}
        &nbsp;
        {{end}}
        {{if .flash.notice}}
        <h3>{{.flash.notice}}</h3>
        &nbsp;
        {{end}}    
        <form method="POST">
            {{.xsrfdata}}
            <div class="form-group col-lg-12">
                <label>Password (min. 6 caratteri)</label>      {{if .Errors.Password}}<b class="error"> ****{{.Errors.Password}}****</b>{{end}}
                <input type="password" name="password" class="form-control"  placeholder="Password">
            </div>
            
            <div class="form-group col-lg-12">
                <label>Ripeti Password</label>      {{if .Errors.Confirm}}<b class="error"><br>****{{.Errors.Confirm}}****</b>{{end}}
                <input type="password" name="password2" class="form-control" placeholder="Password">
            </div>
            <div class="col-sm-12">
                <button type="submit" value="Reset password" class="btn btn-primary pull-right">Reset Password</button>
            </div>	                     
        </form>
        </br>
        <div class="check-help">
            <a href="http://{{.domainname}}">Home</a> 
        </div>                      
    </div>

</div>