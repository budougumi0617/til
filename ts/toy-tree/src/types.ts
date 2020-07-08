export type TODO_any = any;

interface BaseNode<T> {
  type: T;
  name: string;
}

export interface FileNode extends BaseNode<"file"> {}

export interface DirectoryNode extends BaseNode<"directory"> {
  children: TreeNode[];
}

export type TreeNode = FileNode | DirectoryNode;

export interface Options {
  level: number;
}
