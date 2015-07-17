package dirtraverse

import(
        "io/ioutil"
        "path/filepath"
	"os"
	"fmt"
)

type DirTraverse interface {
        Traverse() error
}

type dirTraverse struct {
	dir string
	log string
}

func New(dir string, log string) DirTraverse {
        return dirTraverse{dir, log}
}

 


func (dt dirTraverse) Traverse() error {
        absDir, err := filepath.Abs(dt.dir)
	fmt.Println(absDir)
        if err != nil {
                return err
        } 
        logFile, err := os.Create(dt.log)
        defer logFile.Close()
        if err != nil {
                return err
        }
        return dt.traverse(absDir, logFile)

}

func (dt dirTraverse) traverse(dir string, log *os.File) error {
        infos, err:= ioutil.ReadDir(dir)
        if err != nil {
                return err
	}
        for _,info := range infos {
                if info.IsDir() {
                        if err := dt.traverse(dt.genFullName(dir, info.Name()), log); err != nil {
				return err
			}
                }else {
                        if _,err := log.WriteString(dt.genFullName(dir, info.Name())+"\n"); err != nil {
				return err
			}
		}
	}
	return nil
                                                             
}


func (dt dirTraverse) genFullName(dir string, name string) string {
        return dir+"/"+name
}




