package tpm
/*
#cgo LDFLAGS: -ltspi
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <unistd.h>
#include<signal.h>
#include<sys/param.h>
#include<sys/stat.h>


#include <tss/tss_error.h>
#include <tss/platform.h>  
#include <tss/tss_defines.h>  
#include <tss/tss_typedef.h>  
#include <tss/tss_structs.h>
#include <tss/tss_error.h>
#include <tss/tss_error_basics.h>  
#include <tss/tspi.h>
#include <trousers/trousers.h>
#include <tss/tpm.h>
#include <trousers/tss.h>

int num =0;
static char hashValue[40];

static char basePath[1000];
    
static char orig[40];
FILE *fp;
static int wait_time = 60;



const char *get_error(TSS_RESULT res)
{
    switch(ERROR_CODE(res))
    {        
        case 0x0001L:				return "Authentication failed";
            
        case TSS_E_BAD_PARAMETER:		return "TSS_E_BAD_PARAMETER";
            
        case TSS_E_INTERNAL_ERROR:		return "an error occurred internal to the TSS";
            
        case TSS_E_PS_KEY_NOTFOUND:		return "TSS_E_PS_KEY_NOTFOUND";
            
        case TSS_E_KEY_ALREADY_REGISTERED:	return "TSS_E_KEY_ALREADY_REGISTERED";
            
	case TSS_E_CANCELED:			return "TSS_E_CANCELED";
	case TSS_E_TIMEOUT:			return "TSS_E_TIMEOUT";
	case TSS_E_OUTOFMEMORY:			return "TSS_E_OUTOFMEMORY";
	case TSS_E_TPM_UNEXPECTED:		return "TSS_E_TPM_UNEXPECTED";
	case TSS_E_COMM_FAILURE:		return "TSS_E_COMM_FAILURE";
	case TSS_E_TPM_UNSUPPORTED_FEATURE:	return "TSS_E_TPM_UNSUPPORTED_FEATURE";


        case TSS_E_INVALID_OBJECT_TYPE:		return "TSS_E_INVALID_OBJECT_TYPE";
            
        case TSS_E_INVALID_OBJECT_INITFLAG:	return "TSS_E_INVALID_OBJECT_INITFLAG";
            
	case TSS_E_INVALID_HANDLE:		return "TSS_E_INVALID_HANDLE";
            
	case TSS_E_NO_CONNECTION:		return "TSS_E_NO_CONNECTION";
            
	case TSS_E_CONNECTION_FAILED:		return "TSS_E_CONNECTION_FAILED";
           
	case TSS_E_CONNECTION_BROKEN:		return "TSS_E_CONNECTION_BROKEN";
                     
        case TSS_E_PS_KEY_EXISTS:		return "TSS_E_PS_KEY_EXISTS"; 
                 
        case TSS_E_INVALID_ATTRIB_FLAG:		return "attribflag is incorrect";
            
        case TSS_E_INVALID_ATTRIB_SUBFLAG:	return "subflag is incorrect";
            
        case TSS_E_INVALID_ATTRIB_DATA:		return "ulAttrib is incorrect";
            
        case TSS_E_KEY_NOT_LOADED:		return "TSS_E_KEY_NOT_LOADED";
              
	
	case TSS_E_KEY_NOT_SET:			return "TSS_E_KEY_NOT_SET";
	case TSS_E_VALIDATION_FAILED:		return "TSS_E_VALIDATION_FAILED";
	case TSS_E_TSP_AUTHREQUIRED:		return "TSS_E_TSP_AUTHREQUIRED";
	case TSS_E_TSP_AUTH2REQUIRED:		return "TSS_E_TSP_AUTH2REQUIRED";
	case TSS_E_TSP_AUTHFAIL:		return "TSS_E_TSP_AUTHFAIL";
	case TSS_E_TSP_AUTH2FAIL:		return "TSS_E_TSP_AUTH2FAIL";
	case TSS_E_KEY_NO_MIGRATION_POLICY:	return "TSS_E_KEY_NO_MIGRATION_POLICY";
	case TSS_E_POLICY_NO_SECRET:		return "TSS_E_POLICY_NO_SECRET";
	case TSS_E_INVALID_OBJ_ACCESS:		return "TSS_E_INVALID_OBJ_ACCESS";
	case TSS_E_INVALID_ENCSCHEME:		return "TSS_E_INVALID_ENCSCHEME";
	case TSS_E_INVALID_SIGSCHEME:		return "TSS_E_INVALID_SIGSCHEME";
	case TSS_E_ENC_INVALID_LENGTH:		return "TSS_E_ENC_INVALID_LENGTH";
	case TSS_E_ENC_NO_DATA:			return "TSS_E_ENC_NO_DATA";
	case TSS_E_ENC_INVALID_TYPE:		return "TSS_E_ENC_INVALID_TYPE";
	case TSS_E_INVALID_KEYUSAGE:		return "TSS_E_INVALID_KEYUSAGE";
	case TSS_E_VERIFICATION_FAILED:		return "TSS_E_VERIFICATION_FAILED";
										
    
        case TSS_SUCCESS:			return "success";
            
	
        default:
            return "unknown error";
    }
}


int Hash_File(BYTE *fp,int flen,char *hashValue)
{
TSS_HCONTEXT  hContext;
TSS_HTPM hTPM;
TSS_RESULT result;
TSS_HHASH hHashOfKey;
UINT32    digestLen;
BYTE      *digest;
int        i;



 
    result =Tspi_Context_Create(&hContext);   
    if(result!=TSS_SUCCESS)
    {
        printf("Tsp_Context_Create ERROR:%s(%04x)\n",get_error(result),result);
        Tspi_Context_Close(hContext);
        return -1;
    }
      
    result= Tspi_Context_Connect(hContext,NULL);      
    if(result!=TSS_SUCCESS)
    {
        printf("Tspi_Context_Connect ERROR:%s(%04x)\n",get_error(result),result);
        Tspi_Context_Close(hContext);
        return -1;
    }
      


      result= Tspi_Context_GetTpmObject(hContext,&hTPM);  
      if(result!=TSS_SUCCESS)
      {
          printf("Tspi_Context_GetTpmObject ERROR:%s(%04x)\n",get_error(result),result);
          Tspi_Context_Close(hContext);
          return -1;
      }





     result=Tspi_Context_CreateObject(hContext,TSS_OBJECT_TYPE_HASH,TSS_HASH_SHA1,&hHashOfKey);
     if(result!=TSS_SUCCESS)
      {
          printf("Tspi_Context_CreateObject ERROR:%s(%04x)\n",get_error(result),result);
          Tspi_Context_Close(hContext);
          return -1;
      }





result=Tspi_Hash_UpdateHashValue(hHashOfKey,flen,fp); 
    if(result!=TSS_SUCCESS)
      {
          printf("Tspi_Hash_UpdateHashValue ERROR:%s(%04x)\n",get_error(result),result);
          Tspi_Context_Close(hContext);
          return -1;
      }

result=Tspi_Hash_GetHashValue(hHashOfKey,&digestLen,&digest);
    if(result!=TSS_SUCCESS)
      {
          printf("Tspi_Hash_GetHashValue ERROR:%s(%04x)\n",get_error(result),result);
          Tspi_Context_Close(hContext);
          return -1;
      }
 
free(fp);




memset(hashValue,0,40);

for(i=0;i<digestLen;i++)
   sprintf(hashValue+i*2,"%02x",digest[i]&0xff);
//printf("\nhash: %s\n",hashValue);

   
    Tspi_Context_FreeMemory(hContext,NULL);
    Tspi_Context_Close(hContext);
return 0;

}






int readFileList(char *basePath)
{
    DIR *dir;
    struct dirent *ptr;
    char base[1000];
    
    char fileRoad[1000];
    num++;
    if ((dir=opendir(basePath)) == NULL)
    {
        perror("Open dir error");
        exit(1);
    }

    while ((ptr=readdir(dir)) != NULL)
    {	
        if(strcmp(ptr->d_name,".")==0 || strcmp(ptr->d_name,"..")==0 ||( strcmp(ptr->d_name,".bash_history")==0&&num==2))    ///current dir OR parrent dir
            continue;
        else if(ptr->d_type == 8||ptr->d_type==10)
	{    ///file

	    memset(fileRoad,'\0',sizeof(fileRoad));
            strcpy(fileRoad,basePath);
	    strcat(fileRoad,"/");
	    strcat(fileRoad,ptr->d_name);
	    printf("d_name:%s\n",fileRoad);
 	    FILE *tmp;
	    
	    if((tmp=fopen(fileRoad,"r"))==NULL)
		printf("open %s error\n",fileRoad);
	    else if(fgetc(tmp)!=EOF)
            {
                fclose(tmp);
                tmp=fopen(fileRoad,"r");
		int flen;
		BYTE *fp;
		fseek(tmp,0L,SEEK_END);
		flen=ftell(tmp);
		fp=(BYTE *)malloc(flen+41);

		fseek(tmp,0L,SEEK_SET);
		//fread(fp,flen,1,tmp);
		fp[flen]=0;
		strcat((char *)fp,hashValue);
                Hash_File(fp,flen,hashValue);       
            }	
	    fclose(tmp);
	}
	else if(ptr->d_type == 4)    ///dir
        {
            memset(base,'\0',sizeof(base));
            strcpy(base,basePath);
            strcat(base,"/");
            strcat(base,ptr->d_name);
            readFileList(base);
	    num--;
        }
    }
    closedir(dir);
    return 1;
}

void init_daemon()
{
	int pid;
	int i;
	pid=fork();
	if(pid<0)    
    	    exit(1);  
	else if(pid>0) 
	    exit(0);
	    
	setsid(); 
	pid=fork();
	if(pid>0)
	    exit(0); 
	else if(pid<0)    
	    exit(1);

	for(i=0;i<NOFILE;i++)
	    close(i);
	//chdir("/home");  
	umask(0);
	return ;
}

int bindData(TSS_UUID hKeyUUID,int dataLen,char *dataToBind,int *encDataSize,char **encData);

int tpmtest()
{
    static int wait_num=0;
    ///get the file list
    memset(basePath,'\0',sizeof(basePath));
  //  strcpy(basePath,"/var/lib/docker/aufs/diff/");
         strcpy(basePath,"/home/");
    strcat(basePath,"0cf29d6c76b428ac1d5e755c58a04a97dcb18403d257244c6b9c90ef645e0b6b");
    printf("the tested dir is : %s\n",basePath);
            if(opendir(basePath)==NULL) printf("dir not exist");
    //deamon start!!
    init_daemon();

    //dir exist?
    while(opendir(basePath)!=NULL)
    {
	    
	    memset(hashValue,'\0',sizeof(hashValue));
	    memset(orig,'\0',sizeof(orig));
	    readFileList(basePath);
	    
	    if(strcmp(orig,hashValue)!=0)
	    {
		fp=fopen("hashtest.log","a");
		if(fp>=0)
		{
		    fprintf(fp,"now hash is :%s \n",hashValue); 
		    fclose(fp);
		}
	     }
	
	 
	    sleep(wait_time);
	   // wait_num++;
    }
    return 0;
}
*/
import "C"

func (tpm *Tpm)tpm_main(){
   C.tpmtest()
}

