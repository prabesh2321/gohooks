# gohooks
A go project with pre-commit hook

Instructions

  - clone the project
  - cd gohooks
  - chmod +x hooks/pre-commit
  - mv hooks/pre-commit .git/hooks/pre-commit
  
To see in action

  - git commit -m "tada"

Output

  - Success : if unit test passes
  - Failure : will not be able to commit the changes
     
