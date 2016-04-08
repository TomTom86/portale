

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>{{.UFirst}} {{.ULast}}</h3>
        </div>
        <div class="panel-body">
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}
            {{if .flash.notice}}
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Operazione completata con successo</strong>
                <hr class="message-inner-separator">
                {{.flash.notice}}.
            </div>
            {{end}}            
            {{if .Errors}}
            {{range $rec := .Errors}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{$rec}}.
            </div>      
            {{end}}
            &nbsp;
            {{end}} 
                <form method="POST">
                    {{ .xsrfdata }}         
                    <!-- NOME --> 
                    <div class="form-group">                
                        <label class="control-label" for="first">Nome</label>                          
                        <input id="first" name="first" type="text" value="{{.UFirst}}" class="form-control input-md"/>
                    </div> 
                    <!-- COGNOME --> 
                    <div class="form-group">
                        <label class="control-label" for="last">Cognome</label>                     
                        <input id="last" name="last" type="text" value="{{.ULast}}" class="form-control input-md"/>
                    </div> 
                    <!-- EMAIL -->            
                    <div class="form-group">
                        <label class="control-label" for="email">Email</label>
                        <input id="email" name="email" type="text" value="{{.UEmail}}" class="form-control input-md"/>
                    </div>  
                    <!-- LVL --> 
                    <div class="form-group">
                        <label class="control-label" for="userlvl">Privilegi Utente:</label>
                            <select id="userlvl" name="userlvl" size="1">
                                {{.Userlvllist}}
                            </select>
                    </div>  
                      <!-- BLOCCO --> 
                    <div class="form-group">
                        <label class="control-label" for="blocco">Utente Bloccato:</label> 
                        {{.Checkbloccato}} 
                    </div>   
                    <!-- PASS -->                           
                    <div class="form-group">
                        <label class="control-label" for="password">New password (must be at least 6 characters):</label>
                            <input id="password" name="password" type="password" placeholder="Password" class="form-control input-md">
                            <span class="help-block">Inserisci una password di almeno 6 caratteri</span>
                    </div>  
                    <!-- PASS2 --> 
                    <div class="form-group">
                        <label class="control-label" for="password2">Confirm new password:</label>                            
                            <input id="password2" name="password2" type="password" class="form-control input-md"/>
                    </div>  
                    <!-- IDKEY --> 
                    <div class="form-group">
                        <label class="control-label" for="idkey">ID key</label>  
                            <input id="idkey" name="idkey" value="{{.IDkey}}" size="35" class="form-control input-md"/>
                    </div>  
                    <!-- REGDATE --> 
                    <div class="form-group">
                        <label class="control-label" for="regdate">Data di registrazione</label> 
                            <input id="regdate" name="regdate" value="{{.RegDate}}" size="35" readonly class="form-control input-md"/>
                    </div>   
                    <!-- RESET KEY -->               
                    <div class="form-group">
                        <label class="control-label" for="resetkey">Reset key</label>  
                            <input id="resetkey" name="resetkey" value="{{.ResetKey}}" size="35" class="form-control input-md"/>
                    </div>  
                    <!-- STRUMENTI --> 
                    <div class="form-group  col-xs-2">
                        <label class="control-label" for="apps">Strumenti:</label> 
                    </div>                        
                        <div class="form-group  col-xs-2">
                            {{.Checkautomezzi}}
                    </div>                            
                            <div class="form-group  col-xs-2">
                            {{.Checkservizi}}
                    </div>  
                    <!-- footer --> 
                    <div class="form-group col-xs-12 ">
                        <a href="http://{{.domainname}}/admin/id!0!id__gte,0" >Indietro</a>
                        <input type="submit" value="Update" />
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>